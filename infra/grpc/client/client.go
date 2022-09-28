package main

import (
	"log"

	"google.golang.org/grpc"

	"github.com/FaridehGhani/ompfinex_challenge/middle/proto"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc client error: %v", err)
	}

	_ = proto.NewAsciiArtServiceClient(conn)

	defer conn.Close()
}
