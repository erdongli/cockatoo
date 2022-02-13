package conn

import (
	pb "github.com/erdongli/cockatoo/api"
)

type Connection struct {
	stream pb.Gateway_ConnectServer
	Stat   *Stat
}

func NewConnection(stream pb.Gateway_ConnectServer) *Connection {
	return &Connection{
		stream: stream,
		Stat:   newStat(),
	}
}

func (s *Connection) Close() error {
	return nil
}

func (s *Connection) Send(packet *pb.Packet) error {
	s.Stat.incNumSent()
	return s.stream.Send(packet)
}

func (s *Connection) Recv() (*pb.Packet, error) {
	s.Stat.incNumReceived()
	return s.stream.Recv()
}
