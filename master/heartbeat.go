package master

import (
	"github.com/gobench-io/gobench/pb"
)

// Ping implements ping grpc handler
func (m *Master) Ping(stream pb.Master_PingServer) error {
	return nil
}
