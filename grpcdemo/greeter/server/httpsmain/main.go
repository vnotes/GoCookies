package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/vnotes/gocookies/grpcdemo/greeter/greet"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

type server struct {
	*pb.UnimplementedGreeterServer
}

func (*server) SayHi(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{Message: "Hi " + in.Name}, nil
}

const (
	host = "localhost"
	port = 8069
)

func main() {
	certFile := testdata.Path("server1.pem")
	keyFile := testdata.Path("server1.key")
	caFile := testdata.Path("ca.pem")
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to create credentials")
	}
	credsSrv := credentials.NewServerTLSFromCert(&cert)
	credsCli, err := credentials.NewClientTLSFromFile(caFile, "x.test.youtube.com")
	if err != nil {
		log.Fatalf("failed to create client credentials %s", err)
	}
	opts := []grpc.ServerOption{grpc.Creds(credsSrv)}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(credsCli)}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(grpcServer, &server{})

	ctx := context.Background()
	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	err = pb.RegisterGreeterHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("%s:%d", host, port), dopts)
	if err != nil {
		log.Fatalf("failed to register srv %s", err)
		return
	}
	mux.Handle("/", gwmux)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			NextProtos:   []string{http2.NextProtoTLS},
		},
	}
	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))
	if err != nil {
		log.Fatalf("listen server err %s", err)
	}
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
