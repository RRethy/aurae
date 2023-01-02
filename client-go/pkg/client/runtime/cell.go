package runtime

import (
	"google.golang.org/grpc"

	cellv0 "github.com/aurae-runtime/aurae/client-go/pkg/api/runtime/cell/v0"
)

type CellInterface interface {
	V0() cellv0.CellServiceClient
}

var _ = CellInterface((*cell)(nil))

type cell struct {
	conn *grpc.ClientConn
}

func (c *cell) V0() cellv0.CellServiceClient {
	return cellv0.NewCellServiceClient(c.conn)
}
