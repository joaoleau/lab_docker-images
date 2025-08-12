package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/joaoleau/grpc-ping-server/ping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	PORT string
)

type pingServer struct {
    pb.UnimplementedPingServiceServer
}

func (s *pingServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
    log.Println("Recebido ping")
    return &pb.PingResponse{
        Message: "Pong",
        Code:    200,
    }, nil
}

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = ":50051"
	}
}

func main() {
    lis, err := net.Listen("tcp", PORT)
    if err != nil {
        log.Fatalf("Erro ao abrir porta: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterPingServiceServer(grpcServer, &pingServer{})
	reflection.Register(grpcServer)

    log.Printf("Servidor gRPC rodando na porta %s...", PORT)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Erro ao rodar servidor: %v", err)
    }
}
