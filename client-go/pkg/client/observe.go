package client

import (
	"google.golang.org/grpc"

	observev0 "github.com/aurae-runtime/aurae/client-go/pkg/api/observe/v0"
)

type ObserveInterface interface {
	V0() observev0.ObserveServiceClient
}

var _ = ObserveInterface((*observe)(nil))

type observe struct {
	conn *grpc.ClientConn
}

func (o *observe) V0() observev0.ObserveServiceClient {
	return observev0.NewObserveServiceClient(o.conn)
}
