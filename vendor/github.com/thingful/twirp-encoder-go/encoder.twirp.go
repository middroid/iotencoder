// Code generated by protoc-gen-twirp v5.5.0, DO NOT EDIT.
// source: encoder.proto

/*
Package encoder is a generated twirp stub package.
This code was generated with github.com/twitchtv/twirp/protoc-gen-twirp v5.5.0.

It is generated from these files:
	encoder.proto
*/
package encoder

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Imports only used by utility functions:
import io "io"
import strconv "strconv"
import json "encoding/json"
import url "net/url"

// =================
// Encoder Interface
// =================

// Encoder is the basic interface proposed for the stream encoder component for
// DECODE. It currently just exposes two methods which allow for encoded streams
// to be created and destroyed. Creating a stream means setting up a
// subscription to an MQTT broker such that we start receiving events for a
// specific device. These events are then encrypted using the supplied
// credentials, and then written upstream to our encrypted datastore. Once a
// stream has been created it continues running indefinitely until receiving a
// call to delete the stream.
//
// Later iterations of this service will implement filtering and aggregation
// operations on the stream, but for now all data is simply passed through to
// the datastore.
type Encoder interface {
	// CreateStream sets up a new encoded stream for the encoder. Here we
	// subscribe to the specified MQTT topic, save the encryption keys, and start
	// listening for events. On receiving incoming messages via the MQTT broker,
	// we encrypt the contents using Zenroom and then write the encrypted data to
	// the configured datastore.
	CreateStream(context.Context, *CreateStreamRequest) (*CreateStreamResponse, error)

	// DeleteStream is called to remove the configuration for an encoded data
	// stream. This means deleting the MQTT subscription and removing all saved
	// credentials.
	DeleteStream(context.Context, *DeleteStreamRequest) (*DeleteStreamResponse, error)
}

// =======================
// Encoder Protobuf Client
// =======================

type encoderProtobufClient struct {
	client HTTPClient
	urls   [2]string
}

// NewEncoderProtobufClient creates a Protobuf client that implements the Encoder interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewEncoderProtobufClient(addr string, client HTTPClient) Encoder {
	prefix := urlBase(addr) + EncoderPathPrefix
	urls := [2]string{
		prefix + "CreateStream",
		prefix + "DeleteStream",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &encoderProtobufClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &encoderProtobufClient{
		client: client,
		urls:   urls,
	}
}

func (c *encoderProtobufClient) CreateStream(ctx context.Context, in *CreateStreamRequest) (*CreateStreamResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "decode.iot.encoder")
	ctx = ctxsetters.WithServiceName(ctx, "Encoder")
	ctx = ctxsetters.WithMethodName(ctx, "CreateStream")
	out := new(CreateStreamResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encoderProtobufClient) DeleteStream(ctx context.Context, in *DeleteStreamRequest) (*DeleteStreamResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "decode.iot.encoder")
	ctx = ctxsetters.WithServiceName(ctx, "Encoder")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteStream")
	out := new(DeleteStreamResponse)
	err := doProtobufRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ===================
// Encoder JSON Client
// ===================

type encoderJSONClient struct {
	client HTTPClient
	urls   [2]string
}

// NewEncoderJSONClient creates a JSON client that implements the Encoder interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewEncoderJSONClient(addr string, client HTTPClient) Encoder {
	prefix := urlBase(addr) + EncoderPathPrefix
	urls := [2]string{
		prefix + "CreateStream",
		prefix + "DeleteStream",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &encoderJSONClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &encoderJSONClient{
		client: client,
		urls:   urls,
	}
}

func (c *encoderJSONClient) CreateStream(ctx context.Context, in *CreateStreamRequest) (*CreateStreamResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "decode.iot.encoder")
	ctx = ctxsetters.WithServiceName(ctx, "Encoder")
	ctx = ctxsetters.WithMethodName(ctx, "CreateStream")
	out := new(CreateStreamResponse)
	err := doJSONRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encoderJSONClient) DeleteStream(ctx context.Context, in *DeleteStreamRequest) (*DeleteStreamResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "decode.iot.encoder")
	ctx = ctxsetters.WithServiceName(ctx, "Encoder")
	ctx = ctxsetters.WithMethodName(ctx, "DeleteStream")
	out := new(DeleteStreamResponse)
	err := doJSONRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ======================
// Encoder Server Handler
// ======================

type encoderServer struct {
	Encoder
	hooks *twirp.ServerHooks
}

func NewEncoderServer(svc Encoder, hooks *twirp.ServerHooks) TwirpServer {
	return &encoderServer{
		Encoder: svc,
		hooks:   hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *encoderServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// EncoderPathPrefix is used for all URL paths on a twirp Encoder server.
// Requests are always: POST EncoderPathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const EncoderPathPrefix = "/twirp/decode.iot.encoder.Encoder/"

func (s *encoderServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "decode.iot.encoder")
	ctx = ctxsetters.WithServiceName(ctx, "Encoder")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/twirp/decode.iot.encoder.Encoder/CreateStream":
		s.serveCreateStream(ctx, resp, req)
		return
	case "/twirp/decode.iot.encoder.Encoder/DeleteStream":
		s.serveDeleteStream(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *encoderServer) serveCreateStream(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateStreamJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateStreamProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *encoderServer) serveCreateStreamJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateStream")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(CreateStreamRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *CreateStreamResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.Encoder.CreateStream(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateStreamResponse and nil error while calling CreateStream. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	respBytes := buf.Bytes()
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *encoderServer) serveCreateStreamProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateStream")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(CreateStreamRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *CreateStreamResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.Encoder.CreateStream(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateStreamResponse and nil error while calling CreateStream. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *encoderServer) serveDeleteStream(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveDeleteStreamJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveDeleteStreamProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *encoderServer) serveDeleteStreamJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteStream")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(DeleteStreamRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *DeleteStreamResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.Encoder.DeleteStream(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *DeleteStreamResponse and nil error while calling DeleteStream. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	respBytes := buf.Bytes()
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *encoderServer) serveDeleteStreamProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "DeleteStream")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(DeleteStreamRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *DeleteStreamResponse
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.Encoder.DeleteStream(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *DeleteStreamResponse and nil error while calling DeleteStream. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *encoderServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 0
}

func (s *encoderServer) ProtocGenTwirpVersion() string {
	return "v5.5.0"
}

// =====
// Utils
// =====

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
//
// HTTPClient implementations should not follow redirects. Redirects are
// automatically disabled if *(net/http).Client is passed to client
// constructors. See the withoutRedirects function in this file for more
// details.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwirpServer is the interface generated server structs will support: they're
// HTTP handlers with additional methods for accessing metadata about the
// service. Those accessors are a low-level API for building reflection tools.
// Most people can think of TwirpServers as just http.Handlers.
type TwirpServer interface {
	http.Handler
	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// github.com/golang/protobuf/protoc-gen-go/descriptor.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)
	// ProtocGenTwirpVersion is the semantic version string of the version of
	// twirp used to generate this file.
	ProtocGenTwirpVersion() string
}

// WriteError writes an HTTP response with a valid Twirp error format.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func WriteError(resp http.ResponseWriter, err error) {
	writeError(context.Background(), resp, err, nil)
}

// writeError writes Twirp errors in the response and triggers hooks.
func writeError(ctx context.Context, resp http.ResponseWriter, err error, hooks *twirp.ServerHooks) {
	// Non-twirp errors are wrapped as Internal (default)
	twerr, ok := err.(twirp.Error)
	if !ok {
		twerr = twirp.InternalErrorWith(err)
	}

	statusCode := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
	ctx = ctxsetters.WithStatusCode(ctx, statusCode)
	ctx = callError(ctx, hooks, twerr)

	resp.Header().Set("Content-Type", "application/json") // Error responses are always JSON (instead of protobuf)
	resp.WriteHeader(statusCode)                          // HTTP response status code

	respBody := marshalErrorToJSON(twerr)
	_, writeErr := resp.Write(respBody)
	if writeErr != nil {
		// We have three options here. We could log the error, call the Error
		// hook, or just silently ignore the error.
		//
		// Logging is unacceptable because we don't have a user-controlled
		// logger; writing out to stderr without permission is too rude.
		//
		// Calling the Error hook would confuse users: it would mean the Error
		// hook got called twice for one request, which is likely to lead to
		// duplicated log messages and metrics, no matter how well we document
		// the behavior.
		//
		// Silently ignoring the error is our least-bad option. It's highly
		// likely that the connection is broken and the original 'err' says
		// so anyway.
		_ = writeErr
	}

	callResponseSent(ctx, hooks)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// getCustomHTTPReqHeaders retrieves a copy of any headers that are set in
// a context through the twirp.WithHTTPRequestHeaders function.
// If there are no headers set, or if they have the wrong type, nil is returned.
func getCustomHTTPReqHeaders(ctx context.Context) http.Header {
	header, ok := twirp.HTTPRequestHeaders(ctx)
	if !ok || header == nil {
		return nil
	}
	copied := make(http.Header)
	for k, vv := range header {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}
	return copied
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if customHeader := getCustomHTTPReqHeaders(ctx); customHeader != nil {
		req.Header = customHeader
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Twirp-Version", "v5.5.0")
	return req, nil
}

// JSON serialization for errors
type twerrJSON struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Meta map[string]string `json:"meta,omitempty"`
}

// marshalErrorToJSON returns JSON from a twirp.Error, that can be used as HTTP error response body.
// If serialization fails, it will use a descriptive Internal error instead.
func marshalErrorToJSON(twerr twirp.Error) []byte {
	// make sure that msg is not too large
	msg := twerr.Msg()
	if len(msg) > 1e6 {
		msg = msg[:1e6]
	}

	tj := twerrJSON{
		Code: string(twerr.Code()),
		Msg:  msg,
		Meta: twerr.MetaMap(),
	}

	buf, err := json.Marshal(&tj)
	if err != nil {
		buf = []byte("{\"type\": \"" + twirp.Internal + "\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}

	return buf
}

// errorFromResponse builds a twirp.Error from a non-200 HTTP response.
// If the response has a valid serialized Twirp error, then it's returned.
// If not, the response status code is used to generate a similar twirp
// error. See twirpErrorFromIntermediary for more info on intermediary errors.
func errorFromResponse(resp *http.Response) twirp.Error {
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)

	if isHTTPRedirect(statusCode) {
		// Unexpected redirect: it must be an error from an intermediary.
		// Twirp clients don't follow redirects automatically, Twirp only handles
		// POST requests, redirects should only happen on GET and HEAD requests.
		location := resp.Header.Get("Location")
		msg := fmt.Sprintf("unexpected HTTP status code %d %q received, Location=%q", statusCode, statusText, location)
		return twirpErrorFromIntermediary(statusCode, msg, location)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read server error response body", err)
	}
	var tj twerrJSON
	if err := json.Unmarshal(respBodyBytes, &tj); err != nil {
		// Invalid JSON response; it must be an error from an intermediary.
		msg := fmt.Sprintf("Error from intermediary with HTTP status code %d %q", statusCode, statusText)
		return twirpErrorFromIntermediary(statusCode, msg, string(respBodyBytes))
	}

	errorCode := twirp.ErrorCode(tj.Code)
	if !twirp.IsValidErrorCode(errorCode) {
		msg := "invalid type returned from server error response: " + tj.Code
		return twirp.InternalError(msg)
	}

	twerr := twirp.NewError(errorCode, tj.Msg)
	for k, v := range tj.Meta {
		twerr = twerr.WithMeta(k, v)
	}
	return twerr
}

// twirpErrorFromIntermediary maps HTTP errors from non-twirp sources to twirp errors.
// The mapping is similar to gRPC: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md.
// Returned twirp Errors have some additional metadata for inspection.
func twirpErrorFromIntermediary(status int, msg string, bodyOrLocation string) twirp.Error {
	var code twirp.ErrorCode
	if isHTTPRedirect(status) { // 3xx
		code = twirp.Internal
	} else {
		switch status {
		case 400: // Bad Request
			code = twirp.Internal
		case 401: // Unauthorized
			code = twirp.Unauthenticated
		case 403: // Forbidden
			code = twirp.PermissionDenied
		case 404: // Not Found
			code = twirp.BadRoute
		case 429, 502, 503, 504: // Too Many Requests, Bad Gateway, Service Unavailable, Gateway Timeout
			code = twirp.Unavailable
		default: // All other codes
			code = twirp.Unknown
		}
	}

	twerr := twirp.NewError(code, msg)
	twerr = twerr.WithMeta("http_error_from_intermediary", "true") // to easily know if this error was from intermediary
	twerr = twerr.WithMeta("status_code", strconv.Itoa(status))
	if isHTTPRedirect(status) {
		twerr = twerr.WithMeta("location", bodyOrLocation)
	} else {
		twerr = twerr.WithMeta("body", bodyOrLocation)
	}
	return twerr
}
func isHTTPRedirect(status int) bool {
	return status >= 300 && status <= 399
}

// wrappedError implements the github.com/pkg/errors.Causer interface, allowing errors to be
// examined for their root cause.
type wrappedError struct {
	msg   string
	cause error
}

func wrapErr(err error, msg string) error { return &wrappedError{msg: msg, cause: err} }
func (e *wrappedError) Cause() error      { return e.cause }
func (e *wrappedError) Error() string     { return e.msg + ": " + e.cause.Error() }

// clientError adds consistency to errors generated in the client
func clientError(desc string, err error) twirp.Error {
	return twirp.InternalErrorWith(wrapErr(err, desc))
}

// badRouteError is used when the twirp server cannot route a request
func badRouteError(msg string, method, url string) twirp.Error {
	err := twirp.NewError(twirp.BadRoute, msg)
	err = err.WithMeta("twirp_invalid_route", method+" "+url)
	return err
}

// The standard library will, by default, redirect requests (including POSTs) if it gets a 302 or
// 303 response, and also 301s in go1.8. It redirects by making a second request, changing the
// method to GET and removing the body. This produces very confusing error messages, so instead we
// set a redirect policy that always errors. This stops Go from executing the redirect.
//
// We have to be a little careful in case the user-provided http.Client has its own CheckRedirect
// policy - if so, we'll run through that policy first.
//
// Because this requires modifying the http.Client, we make a new copy of the client and return it.
func withoutRedirects(in *http.Client) *http.Client {
	copy := *in
	copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if in.CheckRedirect != nil {
			// Run the input's redirect if it exists, in case it has side effects, but ignore any error it
			// returns, since we want to use ErrUseLastResponse.
			err := in.CheckRedirect(req, via)
			_ = err // Silly, but this makes sure generated code passes errcheck -blank, which some people use.
		}
		return http.ErrUseLastResponse
	}
	return &copy
}

// doProtobufRequest is common code to make a request to the remote twirp service.
func doProtobufRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return clientError("failed to marshal proto request", err)
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, reqBody, "application/protobuf")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("failed to do request", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read response body", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return clientError("failed to unmarshal proto response", err)
	}
	return nil
}

// doJSONRequest is common code to make a request to the remote twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out proto.Message) (err error) {
	reqBody := bytes.NewBuffer(nil)
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(reqBody, in); err != nil {
		return clientError("failed to marshal json request", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, reqBody, "application/json")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("failed to do request", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(resp.Body, out); err != nil {
		return clientError("failed to unmarshal json response", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}
	return nil
}

// Call twirp.ServerHooks.RequestReceived if the hook is available
func callRequestReceived(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestReceived == nil {
		return ctx, nil
	}
	return h.RequestReceived(ctx)
}

// Call twirp.ServerHooks.RequestRouted if the hook is available
func callRequestRouted(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestRouted == nil {
		return ctx, nil
	}
	return h.RequestRouted(ctx)
}

// Call twirp.ServerHooks.ResponsePrepared if the hook is available
func callResponsePrepared(ctx context.Context, h *twirp.ServerHooks) context.Context {
	if h == nil || h.ResponsePrepared == nil {
		return ctx
	}
	return h.ResponsePrepared(ctx)
}

// Call twirp.ServerHooks.ResponseSent if the hook is available
func callResponseSent(ctx context.Context, h *twirp.ServerHooks) {
	if h == nil || h.ResponseSent == nil {
		return
	}
	h.ResponseSent(ctx)
}

// Call twirp.ServerHooks.Error if the hook is available
func callError(ctx context.Context, h *twirp.ServerHooks, err twirp.Error) context.Context {
	if h == nil || h.Error == nil {
		return ctx
	}
	return h.Error(ctx, err)
}

var twirpFileDescriptor0 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x76, 0xeb, 0xd8, 0x93, 0xb6, 0xb2, 0xb6, 0x11, 0xb2, 0x02, 0x48, 0x21, 0x17, 0x7c,
	0xb2, 0x4a, 0xb9, 0xc0, 0xb1, 0xa5, 0x51, 0x49, 0x13, 0xec, 0xb2, 0x6d, 0x8a, 0xc4, 0xc5, 0x72,
	0xec, 0x11, 0x5a, 0xd5, 0xf1, 0x1a, 0x7b, 0x5d, 0x91, 0xbf, 0xe0, 0xcf, 0xf8, 0x0a, 0xfe, 0x03,
	0x79, 0xed, 0x98, 0x04, 0x22, 0x11, 0xb8, 0x79, 0xde, 0x9b, 0x79, 0xfb, 0x66, 0xc6, 0xbb, 0x70,
	0x88, 0x69, 0xc4, 0x63, 0xcc, 0xdd, 0x2c, 0xe7, 0x82, 0x13, 0x12, 0x63, 0x15, 0xba, 0x8c, 0x0b,
	0xb7, 0x61, 0x86, 0xdf, 0x74, 0x38, 0x7e, 0x9b, 0x63, 0x28, 0xf0, 0x46, 0xe4, 0x18, 0x2e, 0x28,
	0x7e, 0x29, 0xb1, 0x10, 0xe4, 0x39, 0x1c, 0xc4, 0xf8, 0xc0, 0x22, 0x0c, 0x04, 0xbf, 0xc7, 0xd4,
	0x56, 0x06, 0x8a, 0x63, 0xd2, 0x6e, 0x8d, 0xdd, 0x56, 0xd0, 0x5a, 0x4a, 0x12, 0xce, 0x31, 0xb1,
	0xcd, 0xf5, 0x94, 0x69, 0x05, 0x55, 0x29, 0x11, 0x5f, 0x2c, 0xca, 0x94, 0x89, 0x65, 0xc0, 0x62,
	0xdb, 0xa8, 0x53, 0x5a, 0x6c, 0x1c, 0x93, 0x13, 0xe8, 0xe5, 0x18, 0xb1, 0x8c, 0x61, 0x2a, 0x82,
	0xac, 0x9c, 0x27, 0x2c, 0x0a, 0xee, 0x71, 0x69, 0x6b, 0x32, 0x95, 0xb4, 0xdc, 0xb5, 0xa4, 0x26,
	0xb8, 0x24, 0x53, 0x30, 0x12, 0x1e, 0x85, 0x82, 0xf1, 0xd4, 0xde, 0x1f, 0x28, 0x4e, 0xf7, 0xf4,
	0xc4, 0xfd, 0xb3, 0x33, 0x77, 0x4b, 0x57, 0xee, 0xb4, 0xa9, 0xa3, 0xad, 0x42, 0xa5, 0x86, 0x5f,
	0x33, 0x5e, 0x94, 0x39, 0xda, 0xfa, 0x40, 0x71, 0x8e, 0x76, 0x57, 0x1b, 0x35, 0x75, 0xb4, 0x55,
	0x20, 0x1f, 0x00, 0x78, 0x86, 0xb9, 0x94, 0x2e, 0xec, 0xce, 0x40, 0x73, 0xba, 0xa7, 0x2f, 0x77,
	0xd5, 0xf3, 0x57, 0x95, 0x74, 0x4d, 0xa4, 0x7f, 0x01, 0xc6, 0xca, 0x36, 0x79, 0x0a, 0x66, 0xc2,
	0xd3, 0xcf, 0x4c, 0x94, 0x31, 0xca, 0x95, 0x28, 0xf4, 0x17, 0x40, 0xfa, 0x60, 0x24, 0xa1, 0xa8,
	0x49, 0x55, 0x92, 0x6d, 0xdc, 0xff, 0xa1, 0x80, 0xd9, 0xea, 0x93, 0x27, 0x60, 0x16, 0x98, 0x16,
	0x3c, 0xaf, 0x96, 0x52, 0xe9, 0x1c, 0x52, 0xa3, 0x06, 0xc6, 0x31, 0xb9, 0x06, 0x3d, 0x8c, 0xe4,
	0x74, 0x55, 0x39, 0x8f, 0xd7, 0xff, 0xec, 0xdf, 0x3d, 0x93, 0xf5, 0xb4, 0xd1, 0x21, 0x04, 0xf6,
	0xe6, 0x2c, 0x2d, 0x6c, 0x6d, 0xa0, 0x39, 0x0a, 0x95, 0xdf, 0x95, 0x59, 0x96, 0x0a, 0xcc, 0x1f,
	0xc2, 0xc4, 0xde, 0xab, 0x1d, 0xac, 0xe2, 0xe1, 0x1b, 0xd0, 0x6b, 0x05, 0xd2, 0x85, 0xce, 0xcc,
	0x9b, 0x78, 0xfe, 0x47, 0xcf, 0x7a, 0x44, 0x4c, 0xd8, 0xbf, 0x79, 0x77, 0x46, 0x47, 0x96, 0x42,
	0x3a, 0xa0, 0x9d, 0x8f, 0x3d, 0x4b, 0x25, 0x47, 0x00, 0xef, 0xfd, 0xbb, 0xb1, 0x77, 0x19, 0x9c,
	0xdd, 0x5d, 0x5a, 0xda, 0xf0, 0x04, 0x8c, 0xd5, 0x5a, 0x36, 0x8b, 0x01, 0xf4, 0xb1, 0x77, 0xe1,
	0xfb, 0xd4, 0x52, 0x2a, 0xc2, 0x9f, 0xdd, 0xca, 0x40, 0xbd, 0xda, 0x33, 0x54, 0x4b, 0xa3, 0x66,
	0xc6, 0x13, 0x16, 0x55, 0x3f, 0xe9, 0x70, 0x02, 0xbd, 0xcd, 0xee, 0x8a, 0x8c, 0xa7, 0x05, 0x92,
	0x67, 0x00, 0x85, 0x44, 0x82, 0xb2, 0x99, 0x9a, 0x49, 0xcd, 0x1a, 0x99, 0xb1, 0x98, 0xf4, 0x60,
	0xbf, 0xbe, 0x2a, 0xaa, 0x64, 0xea, 0x60, 0x78, 0x05, 0xc7, 0x17, 0x98, 0xe0, 0xef, 0xd7, 0xeb,
	0xbf, 0xb4, 0x1e, 0x43, 0x6f, 0x53, 0xab, 0x36, 0x76, 0xfa, 0x5d, 0x81, 0xce, 0xa8, 0xde, 0x0b,
	0x09, 0xe1, 0x60, 0xdd, 0x3c, 0x79, 0xb1, 0xe3, 0xf2, 0xfa, 0xce, 0xdf, 0x13, 0x9b, 0x39, 0x84,
	0x70, 0xb0, 0x6e, 0x63, 0xfb, 0x11, 0x5b, 0x9a, 0xde, 0x7e, 0xc4, 0xb6, 0x8e, 0xce, 0xcd, 0x4f,
	0x9d, 0x86, 0x9f, 0xeb, 0xf2, 0xed, 0x7a, 0xf5, 0x33, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x00, 0x8c,
	0xdf, 0xcc, 0x04, 0x00, 0x00,
}
