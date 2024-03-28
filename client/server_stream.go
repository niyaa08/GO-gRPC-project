package main

import (
	"context"
	pb "github.com/niyaa08/grpc-project/proto"
	"io"
	"log"
)

func CallSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started...")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finalized!!!")
}
