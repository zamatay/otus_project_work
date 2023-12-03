package monitor_v1

import (
	"context"
)

type ServerApi struct {
	UnimplementedMonitorServer
}

func (s *ServerApi) Connect(context.Context, *RequestConnect) (*OkResponse, error) {
	return &OkResponse{
		Connected: true,
	}, nil
}
