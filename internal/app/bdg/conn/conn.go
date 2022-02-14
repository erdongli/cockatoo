package conn

import (
	"io"
	"log"

	pb "github.com/erdongli/cockatoo/api"
)

type Connection struct {
	Id   string
	Stat *Stat

	stream pb.Gateway_ConnectServer
}

func NewConnection(id string, stream pb.Gateway_ConnectServer) *Connection {
	return &Connection{
		Id:     id,
		Stat:   newStat(),
		stream: stream,
	}
}

func (c *Connection) Listen(errc chan error) {
	for {
		packet, err := c.stream.Recv()
		if err == io.EOF {
			log.Printf("connection aborted for id %s", c.Id)
			errc <- nil
			return
		}
		if err != nil {
			log.Printf("failed to receive packet for id %s: %v", c.Id, err)
			errc <- err
			return
		}

		c.Stat.incNumReceived()

		// send down the connection's context if handlePacket ever involves I/O
		go c.handlePacket(packet)
	}
}

func (c *Connection) handlePacket(_ *pb.Packet) error {
	log.Printf("%s - # received: %d", c.Id, c.Stat.NumReceived())
	return nil
}
