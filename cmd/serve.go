// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"net"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/overmike/webterminal/terminal"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Web Terminal Server",
	Long:  `Run Web Terminal Server`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer(serveOptions)
	},
}

type ServerOptions struct {
	BindAddr string
	Insecure bool
	CertFile string
	KeyFile  string
}

type ServeOptions struct {
	WebSocketServer ServerOptions
	GrpcServer      ServerOptions
	Debug           bool
}

var serveOptions ServeOptions

func runServer(serveOptions ServeOptions) {
	mux := runtime.NewServeMux()

	lis, err := net.Listen("tcp", serveOptions.GrpcServer.BindAddr)
	if err != nil {
		logrus.Fatalf("failed to listen on %v, %v", serveOptions.GrpcServer.BindAddr, err)
		return
	}
	s := grpc.NewServer()
	terminal.RegisterTerminalServer(s, &terminal.Service{})
	reflection.Register(s)
	logrus.Info("Starting Web Terminal")
	go func() {
		logrus.Infof("Starting Web Terminal GRPC Server %v", serveOptions.GrpcServer.BindAddr)
		if err := s.Serve(lis); err != nil {
			logrus.Fatalf("failed to serve: %v", err)
			return
		}
	}()

	err = terminal.RegisterTerminalHandlerFromEndpoint(context.Background(), mux, serveOptions.GrpcServer.BindAddr, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		logrus.Fatalf("Failed to register grpc gateway mux: %v", err)
		return
	}
	logrus.Infof("Starting Web Terminal WebSocket Server %v", serveOptions.WebSocketServer.BindAddr)
	box, err := rice.FindBox("../js/dist/")
	if err != nil {
		logrus.Fatalf("can't find rich box %v", err)
		return
	}
	http.Handle("/terminal", wsproxy.WebsocketProxy(mux))
	http.Handle("/", http.FileServer(box.HTTPBox()))
	if serveOptions.WebSocketServer.Insecure {
		err = http.ListenAndServe(serveOptions.WebSocketServer.BindAddr, nil)
		if err != nil {
			logrus.Fatalf("Listen grpc gateway with mux err: %v", err)
		}
	} else {
		err = http.ListenAndServeTLS(serveOptions.WebSocketServer.BindAddr, serveOptions.WebSocketServer.CertFile, serveOptions.WebSocketServer.KeyFile, nil)
		if err != nil {
			logrus.Fatalf("Listen grpc gateway with mux err: %v", err)
		}
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&serveOptions.WebSocketServer.BindAddr, "addr", "a", ":8081", "API/WS Bind Address")
	serveCmd.Flags().StringVarP(&serveOptions.GrpcServer.BindAddr, "grpc-addr", "g", "127.0.0.1:50051", "GRPC Bind Address")
	serveCmd.Flags().BoolVarP(&serveOptions.WebSocketServer.Insecure, "insecure", "i", false, "Disable TLS")
	serveCmd.Flags().BoolVarP(&serveOptions.Debug, "debug", "d", false, "Debug")

	serveCmd.Flags().StringVarP(&serveOptions.WebSocketServer.CertFile, "cert-file", "c", "~/.webterminal/cert.pem", "Certificate file path")
	serveCmd.Flags().StringVarP(&serveOptions.WebSocketServer.KeyFile, "key-file", "k", "~/.webterminal/key.pem", "Key file path")
}
