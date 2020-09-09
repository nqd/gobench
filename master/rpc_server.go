package master

import (
	"io"

	"github.com/gobench-io/gobench/pb"
)

// SendHeartbeat implements grpc interface for master server
func (m *Master) SendHeartbeat(stream pb.Master_SendHeartbeatServer) error {
	var port int32
	for {
		heartbeat, err := stream.Recv()
		if err == nil {
			if port == 0 {
				port = heartbeat.Port
				m.logger.Infow("agent added")
				// todo: trigger the topology update
			}
			continue
		}

		// err != nil
		if port != 0 {
			m.logger.Infow("agent lost")
			// todo: trigger to topology update
		}

		if err == io.EOF {
			return nil
		}

		return err
	}
}
