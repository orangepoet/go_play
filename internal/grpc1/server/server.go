package main

import (
	"context"
	"go_play/internal/grpc1/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

type GreeterService struct {
	v1.UnimplementedGreeterServer
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

func (svc *GreeterService) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("From HealCheck Check")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (svc *GreeterService) Watch(request *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	log.Println("From HealCheck Watch")
	return nil
}

func (svc *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloResponse, error) {
	return &v1.HelloResponse{
		Message: "Hello " + in.Name,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	v1.RegisterGreeterServer(grpcServer, NewGreeterService())

	//grpc_health_v1.RegisterHealthServer(grpcServer, &greeterService)
	//register.NewConsulCli().RegisterService("greeter", "127.0.0.1", 50052)

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
