package main

import (
	"context"
	"log"
	"time"

	pb "github.com/TheMikeKaisen/go_grpc/proto"
)
func CallSayHello(client pb.GreetServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", response.Message)

}