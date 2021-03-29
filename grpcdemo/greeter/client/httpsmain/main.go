package main

import (
	"context"
	"log"

	pb "github.com/vnotes/gocookies/grpcdemo/greeter/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

func main() {
	caFile := testdata.Path("ca.pem")
	creds, err := credentials.NewClientTLSFromFile(caFile, "x.test.youtube.com")
	if err != nil {
		log.Fatalf("failed to create tls %s", err)
	}
	conn, err := grpc.Dial("localhost:8069", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("connect error %s", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx := context.Background()
	reply, err := c.SayHi(ctx, &pb.GreetRequest{Name: "CTMD"})
	if err != nil {
		log.Fatalf("say hi error %s", err)
	}
	log.Printf("greet %v", reply)
}
