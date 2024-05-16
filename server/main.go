package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/mukappalambda/my-trader/gen/message/v1"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", ":50051", "server address")
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server{})
	log.Printf("server listening at %v", ln.Addr())
	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

func (s *server) PutMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("received: topic: %q, message: %q, created_at: %q", in.GetTopic(), in.GetMessage(), in.GetCreatedAt())
	return &pb.MessageResponse{
		Message: "hello" + in.GetMessage(),
	}, nil
}
