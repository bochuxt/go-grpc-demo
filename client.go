package main

import (
	"demo/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)

	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from client-baohulu!",
	}
	//message.Body="hello from client"

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf(" error from client: %s", err)
	}
	log.Printf(" response from Server: %s", response.Body)
}
