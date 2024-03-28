package main

import (
	"context"
	pb "github.com/niyaa08/grpc-project/proto"
	"log"
	"time"
)

func callSayHi(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("count not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
