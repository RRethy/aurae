package runtime

import (
	"google.golang.org/grpc"
)

type Interface interface {
	Cell() CellInterface
	Pod() PodInterface
	Spawn() SpawnInterface
	Instance() InstanceInterface

	Close() error
}

func NewClient(conn *grpc.ClientConn) Interface {
	return &client{
		conn: conn,
	}
}

var _ = Interface((*client)(nil))

type client struct {
	conn *grpc.ClientConn
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Cell() CellInterface {
	return &cell{conn: c.conn}
}

func (c *client) Pod() PodInterface {
	return &pod{conn: c.conn}
}

func (c *client) Spawn() SpawnInterface {
	return &spawn{conn: c.conn}
}

func (c *client) Instance() InstanceInterface {
	return &instance{conn: c.conn}
}
