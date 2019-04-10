// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: terminal.proto

/*
Package terminal is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package terminal

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_Terminal_Session_0(ctx context.Context, marshaler runtime.Marshaler, client TerminalClient, req *http.Request, pathParams map[string]string) (Terminal_SessionClient, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.Session(ctx)
	if err != nil {
		grpclog.Infof("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, berr
	}
	dec := marshaler.NewDecoder(newReader())
	handleSend := func() error {
		var protoReq SessionRequest
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Infof("Failed to decode request: %v", err)
			return err
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Infof("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	if err := handleSend(); err != nil {
		if cerr := stream.CloseSend(); cerr != nil {
			grpclog.Infof("Failed to terminate client stream: %v", cerr)
		}
		if err == io.EOF {
			return stream, metadata, nil
		}
		return nil, metadata, err
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Infof("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Infof("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

// RegisterTerminalHandlerFromEndpoint is same as RegisterTerminalHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTerminalHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterTerminalHandler(ctx, mux, conn)
}

// RegisterTerminalHandler registers the http handlers for service Terminal to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTerminalHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTerminalHandlerClient(ctx, mux, NewTerminalClient(conn))
}

// RegisterTerminalHandlerClient registers the http handlers for service Terminal
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "TerminalClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TerminalClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TerminalClient" to call the correct interceptors.
func RegisterTerminalHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TerminalClient) error {

	mux.Handle("GET", pattern_Terminal_Session_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Terminal_Session_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Terminal_Session_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Terminal_Session_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"terminal"}, ""))
)

var (
	forward_Terminal_Session_0 = runtime.ForwardResponseStream
)
