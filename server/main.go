package main

import (
	"log"
	"net"
	pb "github.com/TheMikeKaisen/go_grpc/proto"
	"google.golang.org/grpc"
)


const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	// starting a gprc server
	grpcServer := grpc.NewServer() 
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start : %v", err)
	}

/*
net.Listen("tcp", port) → Just reserves port 8080 for listening.
grpc.NewServer() → Creates a gRPC server but does not start it.
grpcServer.Serve(lis) → Starts the gRPC server on port 8080 and makes it handle client requests. 
*/



	
}