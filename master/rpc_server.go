package master

import (
	"fmt"
	"io"
	"net"

	"github.com/gobench-io/gobench/pb"
	"google.golang.org/grpc"
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

func (m *Master) ServeGrpc() error {
	portS := fmt.Sprintf(":%d", m.clusterPort)

	m.logger.Infow("grpc server start", "port", portS)

	lis, err := net.Listen("tcp", portS)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterMasterServer(s, m)
	if err := s.Serve(lis); err != nil {
		m.logger.Errorw("grpc failed to serve", "err", err)
		return err
	}
}
