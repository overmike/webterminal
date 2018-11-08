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
		runServer()
	},
}

func runServer() {
	mux := runtime.NewServeMux()

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatalf("failed to listen on %v:%v", port, err)
		return
	}
	s := grpc.NewServer()
	terminal.RegisterTerminalServer(s, &terminal.Service{})
	reflection.Register(s)
	logrus.Info("Starting Web Terminal")
	go func() {
		logrus.Info("Starting Web Terminal GRPC Server")
		if err := s.Serve(lis); err != nil {
			logrus.Fatalf("failed to serve: %v", err)
			return
		}
	}()

	err = terminal.RegisterTerminalHandlerFromEndpoint(context.Background(), mux, ":50051", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		logrus.Fatalf("Failed to register grpc gateway mux: %v", err)
		return
	}
	logrus.Info("Starting Web Terminal WebSocket Server")
	err = http.ListenAndServe(":8081", wsproxy.WebsocketProxy(mux))
	if err != nil {
		logrus.Fatalf("Listen grpc gateway with mux err: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
