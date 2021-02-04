package main

import (
	"context"
	"flag"
	"gitlab.com/tsuchinaga/hello-go-grpc/hellopb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var port, name string
	flag.StringVar(&port, "port", "0", "server port")
	flag.StringVar(&name, "name", "World", "your name")
	flag.Parse()

	conn, err := grpc.DialContext(context.Background(), "localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	cli := hellopb.NewHelloServiceClient(conn)
	res, err := cli.Hello(context.Background(), &hellopb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", res)
}
