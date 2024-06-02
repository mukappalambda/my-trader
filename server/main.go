package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5"
	pb "github.com/mukappalambda/my-trader/gen/message/v1"
	"github.com/mukappalambda/my-trader/internal/models/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var port = flag.Int("port", 50051, "server port")

type server struct {
	pb.UnimplementedMessageServiceServer
	queries *messages.Queries
}

type App struct {
	*server
	connString string
}

func main() {
	app := NewApp()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func NewApp() *App {
	connString := os.Getenv("DATABASE_URL")
	return &App{connString: connString}
}

func (a *App) Run() error {
	flag.Parse()
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, a.connString)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer conn.Close(ctx)
	queries := messages.New(conn)
	a.server = &server{queries: queries}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		return fmt.Errorf("failed to listen: %q", err)
	}
	defer ln.Close()
	healthcheck := health.NewServer()
	s := grpc.NewServer()

	healthgrpc.RegisterHealthServer(s, healthcheck)
	pb.RegisterMessageServiceServer(s, a.server)

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
