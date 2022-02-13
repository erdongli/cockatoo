package bdg

import (
	"io"
	"log"
	"time"

	pb "github.com/erdongli/cockatoo/api"
	"github.com/erdongli/cockatoo/internal/app/bdg/conn"
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
	conn := conn.NewConnection(stream)
	for {
		packet, err := conn.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("uri: %s, # sent: %d, # received: %d", packet.Uri, conn.Stat.NumSent(), conn.Stat.NumReceived())
	}
}
