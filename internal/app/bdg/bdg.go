package bdg

import (
	"io"
	"log"
	"time"

	pb "github.com/erdongli/cockatoo/api"
	"github.com/erdongli/cockatoo/internal/app/bdg/conn"
	"google.golang.org/grpc/peer"
)

var (
	timeout = time.Duration(2) * time.Second
)

type BDG struct {
	pb.UnimplementedGatewayServer
}

func NewBDG() *BDG {
	return &BDG{}
}

func (g *BDG) Connect(stream pb.Gateway_ConnectServer) error {
	addr := "unknown"
	if peer, ok := peer.FromContext(stream.Context()); ok {
		addr = peer.Addr.String()
	}
	log.Printf("new connection from %s", addr)

	conn := conn.NewConnection(stream)

	errorc := make(chan error)
	go func() {
		for {
			packet, err := conn.Recv()
			if err == io.EOF {
				log.Printf("connection from %s aborted", addr)
				errorc <- nil
				return
			}
			if err != nil {
				log.Printf("failed to receive packet: %v", err)
				errorc <- err
				return
			}

			log.Printf("uri: %s, # sent: %d, # received: %d", packet.Uri, conn.Stat.NumSent(), conn.Stat.NumReceived())
		}
	}()

	return <-errorc
}
