package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/FaridehGhani/ompfinex_challenge/middle/proto"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("grpc server error: %v", err)
	}

	var opt []grpc.ServerOption
	srv := grpc.NewServer(opt...)

	var asciiArtServer proto.AsciiArtServiceServer
	proto.RegisterAsciiArtServiceServer(srv, asciiArtServer)

	err = srv.Serve(listen)
	if err != nil {
		log.Fatalf("grpc serving connection error: %v", err)
	}
}
