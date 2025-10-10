package main

import (
	"log"
	"net"
	pb "server/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//https://buf.build/docs/cli/installation/

func main() {
	var us UserService
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Println(err)
		return
	}

	// Create a gRPC server variable of type *grpc.Server
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &us)

	reflection.Register(s)

	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}

}
