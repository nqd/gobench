package master

import (
	"fmt"
	"net"

	"github.com/gobench-io/gobench/pb"
	"google.golang.org/grpc"
)

func (m *Master) startGrpcServer() error {
	addr := fmt.Sprintf(":%d", m.clusterPort)

	m.logger.Infow("start gRPC server at", "address", addr)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	pb.RegisterMasterServer(s, m)

	go s.Serve(l)

	return nil
}
