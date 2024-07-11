package main

import (
	"context"
	"go_play/internal/grpc1/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

// const target = "consul://127.0.0.1:8500/greeter"
const target = "localhost:8888"

func main() {
	//consulRegister := &register.ConsulRegister{}
	//consulRegister.Register()

	ctx0, cancel0 := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel0()

	conn, err := grpc.DialContext(ctx0, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithUnaryInterceptor(unaryInterceptor))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := v1.NewGreeterClient(conn)
	ctx1, cancel1 := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel1()

	resp, err := client.SayHello(ctx1, &v1.HelloRequest{Name: "chengzhi..."})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Message)
}

func unaryInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	startTime := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	duration := time.Since(startTime)
	log.Printf("Method: %s, Duration: %s, Error: %v", method, duration, err)
	if err != nil {
		// 可以在这里添加错误处理逻辑
		st, _ := status.FromError(err)
		if st.Code() == codes.ResourceExhausted {
			log.Printf("Request was rejected with ResourceExhausted error")
		}
	}
	return err
}
