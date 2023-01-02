package runtime

import (
	"google.golang.org/grpc"

	spawnv0 "github.com/aurae-runtime/aurae/client-go/pkg/api/runtime/spawn/v0"
)

type SpawnInterface interface {
	V0() spawnv0.SpawnServiceClient
}

var _ = SpawnInterface((*spawn)(nil))

type spawn struct {
	conn *grpc.ClientConn
}

func (s *spawn) V0() spawnv0.SpawnServiceClient {
	return spawnv0.NewSpawnServiceClient(s.conn)
}
