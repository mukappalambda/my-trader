package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/jackc/pgx/v5"
	"github.com/mukappalambda/my-trader/internal/adapters/database/postgres/messages"
	pb "github.com/mukappalambda/my-trader/internal/adapters/grpc/message/v1"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	pb.UnimplementedMessageServiceServer
	queries *messages.Queries
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

type App struct {
	grpcServer *GrpcServer
	connString string
}

func NewApp() *App {
	connString := os.Getenv("DATABASE_URL")
	return &App{connString: connString}
}
func (a *App) Run(port int) error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, a.connString)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer conn.Close(ctx)
	queries := messages.New(conn)
	a.grpcServer = &GrpcServer{queries: queries}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %q", err)
	}
	defer ln.Close()
	healthcheck := health.NewServer()
	s := grpc.NewServer()

	healthgrpc.RegisterHealthServer(s, healthcheck)
	pb.RegisterMessageServiceServer(s, a.grpcServer)

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
