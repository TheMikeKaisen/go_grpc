package main

import (
	"log"

	pb "github.com/TheMikeKaisen/go_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// UNARY STREAMING
	// CallSayHello(client)

	names := &pb.NamesList{
		Names: []string{"Karthik", "Alice", "Bob"},
	}
	// SERVER STREAMING
	// CallSayHelloServerStream(client, names);

	// Client STREAMING
	// CallSayHelloClientStream(client, names)

	//Bi-Directional Streaming
	CallHelloBidirectionalStream(client, names)

}
