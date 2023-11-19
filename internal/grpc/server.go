package grpc

import (
	"context"
	monitor_v1 "project_work/gen"
)

type serverApi struct {
	monitor_v1.UnimplementedMonitorServer
	monitor Monitor
}

type Monitor interface {
	Connect(ctx context.Context, requestConnect *monitor_v1.RequestConnect)
}

func (s *serverApi) Connect(ctx context.Context, requestConnect *monitor_v1.RequestConnect) {
	//
}
