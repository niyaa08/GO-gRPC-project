package main

import (
	pb "github.com/niyaa08/grpc-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8080"
)

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Niya", "Misha", "Jack"},
	}

	//callSayHi(client)
	//CallSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	CallSayHelloBidirectionalStream(client, names)

}
