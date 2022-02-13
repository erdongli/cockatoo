package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/erdongli/cockatoo/api"
	"github.com/erdongli/cockatoo/internal/app/bdg"
)

var port = flag.Int("port", 50051, "the server port")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	grpcServer := grpc.NewServer()
	pb.RegisterGatewayServer(grpcServer, bdg.NewBDG())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
