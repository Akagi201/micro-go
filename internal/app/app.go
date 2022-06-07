package app

import (
	"context"
	"errors"
	"log"

	apppb "github.com/Akagi201/micro-go/internal/proto/app"
)

type Server struct {
	apppb.UnimplementedMicroAppServer
}

func NewServer() *Server {
	return &Server{}
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *apppb.HelloReq) (*apppb.HelloResp, error) {
	log.Printf("Received: %v", in.GetName())
	if len(in.Name) == 0 {
		return nil, errors.New("name cannot be empty")
	}

	return &apppb.HelloResp{Message: "Hello " + in.Name}, nil
}
