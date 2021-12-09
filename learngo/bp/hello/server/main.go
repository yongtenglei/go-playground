package main

import (
	"context"
	"fmt"
	"net"

	hello_grpc "github.com/hello_grpc/pb"
	"google.golang.org/grpc"
)

type Server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *Server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "from server, "}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &Server{})
	s.Serve(l)

}
