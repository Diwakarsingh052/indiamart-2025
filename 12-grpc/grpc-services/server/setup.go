package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

//https://buf.build/docs/cli/installation/

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Println(err)
		return
	}

	// Create a gRPC server variable of type *grpc.Server
	s := grpc.NewServer()

	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}

}
