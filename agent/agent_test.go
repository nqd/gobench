package agent

import (
	"context"
	"testing"

	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/pb"
	"github.com/stretchr/testify/assert"
)

// nop metric logger
type nopLog struct{}

func (n *nopLog) FindCreateGroup(context.Context, *pb.FCGroupReq) (*pb.FCGroupRes, error) {
	return nil, nil
}
func (n *nopLog) FindCreateGraph(context.Context, *pb.FCGraphReq) (*pb.FCGraphRes, error) {
	return nil, nil
}
func (n *nopLog) FindCreateMetric(context.Context, *pb.FCMetricReq) (*pb.FCMetricRes, error) {
	return nil, nil
}
func (n *nopLog) Histogram(context.Context, *pb.HistogramReq) (*pb.HistogramRes, error) {
	return nil, nil
}
func (n *nopLog) Counter(context.Context, *pb.CounterReq) (*pb.CounterRes, error) {
	return nil, nil
}
func (n *nopLog) Gauge(context.Context, *pb.GaugeReq) (*pb.GaugeRes, error) {
	return nil, nil
}

func newNopMetricLog() *nopLog {
	return &nopLog{}
}

func newAgent(t *testing.T, opts *Options) *Agent {
	logger := logger.NewNopLogger()
	ml := newNopMetricLog()

	a, err := NewAgent(opts, ml, logger)
	assert.Nil(t, err)

	return a
}

func TestNewAgent(t *testing.T) {
	opts := &Options{
		Route:       "localhost:1234",
		ClusterPort: 2345,
	}
	logger := logger.NewNopLogger()
	ml := newNopMetricLog()

	_, err := NewAgent(opts, ml, logger)
	assert.Nil(t, err)
}

func TestStartAgent(t *testing.T) {
	a := newAgent(t, &Options{
		Route:       "localhost:1234",
		ClusterPort: 2345,
	})
	assert.Nil(t, a.StartSocketServer())
	// insert the grpc over tcp here
}
