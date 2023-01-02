package runtime

import (
	"google.golang.org/grpc"

	podv0 "github.com/aurae-runtime/aurae/client-go/pkg/api/runtime/pod/v0"
)

type PodInterface interface {
	V0() podv0.PodServiceClient
}

var _ = PodInterface((*pod)(nil))

type pod struct {
	conn *grpc.ClientConn
}

func (p *pod) V0() podv0.PodServiceClient {
	return podv0.NewPodServiceClient(p.conn)
}
