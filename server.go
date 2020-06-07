package main

import (
	"demo/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf(" failed to listen on port 9000:%v", err)
	}
	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to server over port 9000: %v", err)
	}
	log.Print("server started")
}
