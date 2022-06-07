package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Akagi201/micro-go/internal/app"
	apppb "github.com/Akagi201/micro-go/internal/proto/app"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8327, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := app.NewServer()
	apppb.RegisterMicroAppServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
