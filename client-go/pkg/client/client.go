package client

import (
	"context"

	"google.golang.org/grpc"

	"github.com/aurae-runtime/aurae/client-go/pkg/client/runtime"
)

type Interface interface {
	Observe() ObserveInterface
	Runtime() runtime.Interface

	Close() error
}

func NewClient(ctx context.Context, addr string, dialOpts ...grpc.DialOption) (Interface, error) {
	conn, err := grpc.Dial(addr, dialOpts)
	if err != nil {
		return nil, err
	}

	return &client{
		conn: conn,
	}, nil
}

var _ = Interface((*client)(nil))

type client struct {
	conn *grpc.ClientConn
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Observe() ObserveInterface {
	return &observe{conn: c.conn}
}

func (c *client) Runtime() runtime.Interface {
	return runtime.NewClient(c.conn)
}
