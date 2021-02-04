package main

import (
	"context"
	"fmt"
	"gitlab.com/tsuchinaga/hello-go-grpc/hellopb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("address is", ln.Addr())

	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(ln); err != nil {
		log.Fatalln(err)
	}
}

type server struct {
	hellopb.UnimplementedHelloServiceServer
}

func (s *server) Hello(_ context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	log.Println("get request", req)
	return &hellopb.HelloResponse{Message: fmt.Sprintf("Hello, %s", req.Name)}, nil
}
