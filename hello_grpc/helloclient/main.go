package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-go-course/hello/hellopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client.")

	options := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := hellopb.NewHelloServiceClient(cc)
	request := &hellopb.HelloRequest{Name: "Murali"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Printf("Recieve response => [%v]", resp.Greeting)
}
