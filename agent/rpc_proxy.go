package agent

import (
	"context"

	"github.com/gobench-io/gobench/pb"
)

func (a *Agent) Counter(ctx context.Context, req *pb.CounterReq) (*pb.CounterRes, error) {
	return nil, nil
}

func (a *Agent) Histogram(ctx context.Context, req *pb.HistogramReq) (*pb.HistogramRes, error) {
	return nil, nil
}

func (a *Agent) Histogram(ctx context.Context, req *pb.HistogramReq) (*pb.HistogramRes, error) {
	return nil, nil
}
func (a *Agent) Gauge(ctx context.Context, req *pb.GaugeReq) (*pb.GaugeRes, error) {
	return nil, nil
}
func (a *Agent) FindCreateGroup(ctx context.Context, req *pb.FCGroupReq) (res *pb.FCGroupRes, err error) {
	return nil, nil
}
func (a *Agent) FindCreateGraph(ctx context.Context, req *pb.FCGraphReq) (res *pb.FCGraphRes, err error) {
	return nil, nil
}
func (a *Agent) FindCreateMetric(ctx context.Context, req *pb.FCMetricReq) (res *pb.FCMetricRes, err error) {
	return nil, nil
}
