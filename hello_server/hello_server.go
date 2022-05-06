package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/howDeepTheRabbitHoleGoes/hello_grpc/proto"

	"google.golang.org/grpc"
)

const (
	port = "localhost:50051"
)

type helloServer struct {
	pb.UnimplementedHelloServer
}

func (hs *helloServer) Say(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	fmt.Println("We are before sending")
	response := &pb.Response{
		Message: fmt.Sprintf("Hello %s", request.Name),
	}
	fmt.Println("We are after sending")
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Fail to listen at %s", lis.Addr())
	}
	gs := grpc.NewServer()
	pb.RegisterHelloServer(gs, &helloServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := gs.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
