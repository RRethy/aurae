package runtime

import (
	"google.golang.org/grpc"

	instancev0 "github.com/aurae-runtime/aurae/client-go/pkg/api/runtime/instance/v0"
)

type InstanceInterface interface {
	V0() instancev0.InstanceServiceClient
}

var _ = InstanceInterface((*instance)(nil))

type instance struct {
	conn *grpc.ClientConn
}

func (i *instance) V0() instancev0.InstanceServiceClient {
	return instancev0.NewInstanceServiceClient(i.conn)
}
