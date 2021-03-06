package main

import (
	"context"
	"log"

	pb "github.com/vnotes/gocookies/grpcdemo/greeter/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:8069", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error %s", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "X-A-bC-De", "ojkk")
	reply, err := c.SayHi(ctx, &pb.GreetRequest{Name: "CTMD"})
	if err != nil {
		log.Fatalf("say hi error %s", err)
	}
	log.Printf("greet %v", reply)
}
