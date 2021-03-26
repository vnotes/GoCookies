package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/vnotes/gocookies/grpcdemo/greeter/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:8069", opts)
	if err != nil {
		grpclog.Fatalf("regist srv error %s", err)
	}
	grpclog.Infoln("server on 8070")
	err = http.ListenAndServe(":8070", mux)
	if err != nil {
		grpclog.Fatalf("listen error %s", err)
	}
}
