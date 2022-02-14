package bdg

import (
	"fmt"
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

	registry *conn.Registry
}

func NewBDG() *BDG {
	return &BDG{
		registry: conn.NewRegistry(),
	}
}

func (g *BDG) Connect(stream pb.Gateway_ConnectServer) error {
	// using source IP address as id for now
	peer, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("failed to parse peer information")
	}
	addr := peer.Addr.String()

	conn := conn.NewConnection(addr, stream)
	g.registry.Add(conn)
	defer g.registry.Del(conn)

	errc := make(chan error)
	go conn.Listen(errc)

	return <-errc
}
