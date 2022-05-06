package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/howDeepTheRabbitHoleGoes/hello_grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "2.tcp.ngrok.io:18718", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Did not connect")
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	r, err := c.Say(ctx, &pb.Request{Name: *name})

	if err != nil {
		log.Fatalf("Could not greet %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
