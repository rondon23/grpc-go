package main

import (
	"context"
	"log"

	pb "github.com/rondon23/grpc-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Bruno",
	})

	if err != nil {
		log.Fatal("Cold not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
