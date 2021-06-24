/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
  "context"
  "log"
  "net"

  "google.golang.org/grpc"
  pb "google.golang.org/grpc/examples/helloworld/helloworld"
    "google.golang.org/grpc/reflection"
)

const (
  port = ":31001"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
  pb.UnimplementedTrafficTestApiServer
}


func (s *server) TrafficRequest(ctx context.Context, in *pb.TrafficTestRequest) (*pb.TrafficTestResponse, error) {
  // TODO: Add support for extracting response size and latency injection requested by client.
  // For now just return a 10 bytes hardcoded string in response with 0 latency
  return &pb.TrafficTestResponse{Message: "abcdefghij"}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterTrafficTestApiServer(s, &server{})
  reflection.Register(s)
  log.Printf("server listening at %v", lis.Addr())
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
