package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"

	pb "github.com/erdongli/cockatoo/api"
)

var addr = flag.String("addr", "localhost:50051", "the server address")

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGatewayClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Bi-Directional Gateway")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read from stdin: %v", err)
		}
		cmd = strings.Replace(cmd, "\n", "", -1)

		if cmd == "exit" {
			break
		}

		if err := stream.Send(&pb.Packet{Uri: cmd}); err != nil {
			log.Fatalf("failed to send a packet: %v", err)
		}
	}

	stream.CloseSend()
}
