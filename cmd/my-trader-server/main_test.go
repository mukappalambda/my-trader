package main

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	grpcAdapter "github.com/mukappalambda/my-trader/internal/adapters/grpc"
	pb "github.com/mukappalambda/my-trader/internal/adapters/grpc/message/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMessageService(t *testing.T) {
	var lc net.ListenConfig
	ln, err := lc.Listen(t.Context(), "tcp", ":")
	if err != nil {
		t.Fatalf("failed to listen: %q", err)
	}
	t.Cleanup(func() {
		ln.Close()
	})
	s := grpc.NewServer()
	t.Cleanup(func() {
		s.Stop()
	})
	pb.RegisterMessageServiceServer(s, &grpcAdapter.GrpcServer{})
	go func() {
		log.Println(ln.Addr().String())
		if err := s.Serve(ln); err != nil {
			log.Fatalf("failed to serve: %q\n", err)
		}
	}()

	conn, err := grpc.NewClient(ln.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to create a new client: %q\n", err)
	}
	t.Cleanup(func() {
		conn.Close()
	})
	c := pb.NewMessageServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	t.Cleanup(func() {
		cancel()
	})
	message := "test-message"
	r, err := c.PutMessage(ctx, &pb.MessageRequest{
		Topic:   "test-topic",
		Message: message,
	})
	if err != nil {
		t.Fatalf("failed to call PutMessage: %q\n", err)
	}
	got := r.GetMessage()
	want := "hello" + message

	if got != want {
		t.Fatalf("got: %v; want: %v\n", got, want)
	}
}
