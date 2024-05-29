package main

import (
	"context"
	pb "go_play/api/grpc/v1"
	"go_play/internal/grpc1/register"
	"google.golang.org/grpc"
	"log"
	"time"
)

const target = "consul://127.0.0.1:8500/greeter"

func main() {
	register.Init()
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "chengzhi..."})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Message)
}
