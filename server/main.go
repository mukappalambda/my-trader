package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	pb "github.com/mukappalambda/my-trader/gen/message/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "server port")
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

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return fmt.Errorf("failed to listen: %q", err)
	}
	defer ln.Close()
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, srv)
	reflection.Register(s)
	log.Printf("server listening at %v", ln.Addr())
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := s.Serve(ln); err != nil {
			log.Fatalf("failed to serve: %q", err)
		}
	}()
	<-ctx.Done()
	stop()
	log.Println("server shutting down...")
	s.GracefulStop()
	log.Println("server is down.")
	return nil
}

func (s *server) PutMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("received: topic: %q, message: %q, created_at: %q", in.GetTopic(), in.GetMessage(), in.GetCreatedAt())
	return &pb.MessageResponse{
		Message: "hello" + in.GetMessage(),
	}, nil
}
