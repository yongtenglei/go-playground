package main

import (
	"context"
	"fmt"

	hello_grpc "github.com/hello_grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, e := grpc.Dial("localhost:8888", grpc.WithInsecure())
	defer conn.Close()
	fmt.Println(e)
	client := hello_grpc.NewHelloGRPCClient(conn)
	req, _ := client.SayHi(context.Background(), &hello_grpc.Req{Message: "from client"})
	fmt.Println(req.GetMessage())
}
