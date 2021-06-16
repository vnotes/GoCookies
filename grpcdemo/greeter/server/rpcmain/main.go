package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/vnotes/gocookies/grpcdemo/greeter/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	*pb.UnimplementedGreeterServer
}

func (*server) SayHi(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Printf("got ctx %#v %t\n", md, ok)
	return &pb.GreetResponse{Message: "Hi " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8069")
	if err != nil {
		log.Fatalf("listen network error %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server %s", err)
	}
}
