package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/jackc/pgx/v5"
	"github.com/mukappalambda/my-trader/internal/adapters/database/postgres/messages"
	pb "github.com/mukappalambda/my-trader/internal/adapters/grpc/message/v1"
	"github.com/mukappalambda/my-trader/internal/core/domain/entities"
	"github.com/mukappalambda/my-trader/internal/core/ports"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	pb.UnimplementedMessageServiceServer
	queries *messages.Queries
	conn    *pgx.Conn
}

func NewGrpcServer(dsn string) (*GrpcServer, error) {
	ctx := context.Background()
	return NewGrpcServerWithContext(ctx, dsn)
}

func NewGrpcServerWithContext(ctx context.Context, dsn string) (*GrpcServer, error) {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	queries := messages.New(conn)
	s := &GrpcServer{queries: queries, conn: conn}
	return s, nil
}

var _ ports.MessageService = (*GrpcServer)(nil)

func (s *GrpcServer) PublishMessage(ctx context.Context, msg *entities.Message) (v interface{}, err error) {
	in := &pb.MessageRequest{
		Topic:   msg.Topic,
		Message: msg.Message,
	}
	return s.PutMessage(ctx, in)
}

func (s *GrpcServer) PutMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Printf("received: topic: %q, message: %q, created_at: %q", in.GetTopic(), in.GetMessage(), in.GetCreatedAt())
	_, err := s.queries.CreateMessage(ctx, messages.CreateMessageParams{
		Topic:   in.GetTopic(),
		Message: in.GetMessage(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.MessageResponse{
		Message: "hello" + in.GetMessage(),
	}, nil
}

func (s *GrpcServer) Run(port int) error {
	ctx := context.Background()
	return s.RunWithContext(ctx, port)
}

func (s *GrpcServer) RunWithContext(ctx context.Context, port int) error {
	defer s.conn.Close(ctx)

	var lc net.ListenConfig
	ln, err := lc.Listen(ctx, "tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %q", err)
	}
	defer ln.Close()
	healthcheck := health.NewServer()
	srv := grpc.NewServer()

	healthgrpc.RegisterHealthServer(srv, healthcheck)
	pb.RegisterMessageServiceServer(srv, s)

	reflection.Register(srv)
	log.Printf("server listening at %v", ln.Addr())
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := srv.Serve(ln); err != nil {
			log.Fatalf("failed to serve: %q", err)
		}
	}()
	<-ctx.Done()
	stop()
	log.Println("server shutting down...")
	srv.GracefulStop()
	log.Println("server is down.")
	return nil
}
