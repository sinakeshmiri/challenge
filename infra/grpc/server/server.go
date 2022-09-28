package main

import (
	"github.com/FaridehGhani/ompfinex_challenge/middle/proto"
	"google.golang.org/grpc"
	"log"
	"net"
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
