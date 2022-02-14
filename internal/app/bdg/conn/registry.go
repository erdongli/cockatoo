package conn

import (
	"fmt"
	"log"
	"sync"
)

type Registry struct {
	mutex sync.Mutex
	conns map[string][]*Connection
}

func NewRegistry() *Registry {
	return &Registry{conns: make(map[string][]*Connection)}
}

func (r *Registry) Add(conn *Connection) error {
	log.Printf("adding connection for id %s", conn.Id)

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.conns[conn.Id] = append(r.conns[conn.Id], conn)
	return nil
}

func (r *Registry) Del(conn *Connection) error {
	log.Printf("deleting connection for id %s", conn.Id)

	r.mutex.Lock()
	defer r.mutex.Unlock()

	if len(r.conns[conn.Id]) == 1 {
		delete(r.conns, conn.Id)
		return nil
	}

	for i, c := range r.conns[conn.Id] {
		if c == conn {
			r.conns[conn.Id][i] = r.conns[conn.Id][len(r.conns[conn.Id])-1]
			r.conns[conn.Id] = r.conns[conn.Id][:len(r.conns[conn.Id])-1]
			return nil
		}
	}

	return fmt.Errorf("connection not found for id %s", conn.Id)
}
