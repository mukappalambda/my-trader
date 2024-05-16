package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/mukappalambda/my-trader/gen/message/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr = flag.String("addr", ":50051", "server address")
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func main() {
	if err := run(&server{}); err != nil {
		log.Fatal(err)
	}
}

func run(srv pb.MessageServiceServer) error {
	flag.Parse()

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %q", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, srv)
	reflection.Register(s)
	log.Printf("server listening at %v", ln.Addr())
	if err := s.Serve(ln); err != nil {
		return fmt.Errorf("failed to serve: %q", err)
	}
	return nil
}

func (s *server) PutMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("received: topic: %q, message: %q, created_at: %q", in.GetTopic(), in.GetMessage(), in.GetCreatedAt())
	return &pb.MessageResponse{
		Message: "hello" + in.GetMessage(),
	}, nil
}
