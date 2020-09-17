package agent

import (
	"context"
	"time"

	"github.com/gobench-io/gobench/pb"
	"google.golang.org/grpc"
)

func (a *Agent) Heartbeat() {
	for {
		err := a.doHeartbeat(20 * time.Second)
		if err != nil {
			time.Sleep(30 * time.Second)
		}
	}
}

// ConnectMaster starts the heartbeat to master
func (a *Agent) doHeartbeat(hi time.Duration) error {
	conn, err := grpc.Dial(a.route, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	c := pb.NewMasterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.SendHeartbeat(ctx)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(hi)

	for {
		select {
		case <-ticker.C:
			if err := a.sendOneHeartbeat(stream); err != nil {
				return err
			}
		}
	}
}

func (a *Agent) sendOneHeartbeat(stream pb.Master_SendHeartbeatClient) error {
	a.logger.Infow("send one heartbeat")
	beat := &pb.Heartbeat{
		Port: int32(a.clusterPort),
	}
	err := stream.Send(beat)
	return err
}
