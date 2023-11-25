package monitor_v1

import (
	"project_work/internal/services/sysInfo"
	"time"
)

type ServerApi struct {
	UnimplementedMonitorServer
	SysInfo *sysInfo.SysInfoSrv
}

func NewServerApi(sysInfo *sysInfo.SysInfoSrv) *ServerApi {
	return &ServerApi{SysInfo: sysInfo}
}

func (s *ServerApi) Connect(in *RequestConnect, srv Monitor_ConnectServer) error {
	for {
		data := s.SysInfo.ResultWork.AddItem
		srv.Send()
		time.Sleep(time.Duration(in.EachNSec) * time.Second)
	}
	return nil
}
