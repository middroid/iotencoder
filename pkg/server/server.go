package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	kitlog "github.com/go-kit/kit/log"
	twrpprom "github.com/joneskoo/twirp-serverhook-prometheus"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	datastore "github.com/thingful/twirp-datastore-go"
	encoder "github.com/thingful/twirp-encoder-go"

	"github.com/thingful/iotencoder/pkg/mqtt"
	"github.com/thingful/iotencoder/pkg/postgres"
	"github.com/thingful/iotencoder/pkg/rpc"
	"github.com/thingful/iotencoder/pkg/system"
)

// Server is our top level type, contains all other components, is responsible
// for starting and stopping them in the correct order.
type Server struct {
	srv    *http.Server
	enc    *rpc.Encoder
	db     postgres.DB
	mqtt   mqtt.Client
	logger kitlog.Logger
}

// PulseHandler is the simplest possible handler function - used to expose an
// endpoint which a load balancer can ping to verify that a node is running and
// accepting connections.
func PulseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// NewServer returns a new simple HTTP server.
func NewServer(addr, connStr, encryptionPassword string, logger kitlog.Logger) *Server {
	db := postgres.NewDB(connStr, encryptionPassword, logger)

	ds := datastore.NewDatastoreProtobufClient("http://192.168.1.116:8081", &http.Client{})

	mc := mqtt.NewClient(logger, db, ds)

	enc := rpc.NewEncoder(logger, mc, db)
	hooks := twrpprom.NewServerHooks(nil)

	logger = kitlog.With(logger, "module", "server")
	logger.Log("msg", "creating server")

	twirpHandler := encoder.NewEncoderServer(enc, hooks)

	// multiplex twirp handler into a mux with our other handlers
	mux := http.NewServeMux()
	mux.Handle(encoder.EncoderPathPrefix, twirpHandler)
	mux.HandleFunc("/pulse", PulseHandler)
	mux.Handle("/metrics", promhttp.Handler())

	// create our http.Server instance
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// return the instantiated server
	return &Server{
		srv:    srv,
		enc:    enc,
		db:     db,
		mqtt:   mc,
		logger: kitlog.With(logger, "module", "server"),
	}
}

// Start starts the server running. This is responsible for starting components
// in the correct order, and in addition we attempt to run all up migrations as
// we start.
//
// We also create a channel listening for interrupt signals before gracefully
// shutting down.
func (s *Server) Start() error {
	// start the postgres connection pool
	err := s.db.(system.Component).Start()
	if err != nil {
		return errors.Wrap(err, "failed to start db")
	}

	// start the mqtt client
	err = s.mqtt.(system.Component).Start()
	if err != nil {
		return errors.Wrap(err, "failed to start mqtt client")
	}

	// migrate up the database
	err = s.db.MigrateUp()
	if err != nil {
		return errors.Wrap(err, "failed to migrate the database")
	}

	// start the encoder RPC service
	err = s.enc.Start()
	if err != nil {
		return errors.Wrap(err, "failed to start encoder")
	}

	// add signal handling stuff to shutdown gracefully
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		s.logger.Log("listenAddr", s.srv.Addr, "msg", "starting server")
		if err := s.srv.ListenAndServe(); err != nil {
			s.logger.Log("err", err)
			os.Exit(1)
		}
	}()

	<-stopChan
	return s.Stop()
}

func (s *Server) Stop() error {
	s.logger.Log("msg", "stopping")
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	err := s.enc.Stop()
	if err != nil {
		return err
	}

	err = s.mqtt.(system.Component).Stop()
	if err != nil {
		return err
	}

	err = s.db.(system.Component).Stop()
	if err != nil {
		return err
	}

	return s.srv.Shutdown(ctx)
}
