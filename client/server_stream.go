package main

import (
	"context"
	"io"
	"log"

	pb "github.com/TheMikeKaisen/go_grpc/proto"
)

func CallSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
