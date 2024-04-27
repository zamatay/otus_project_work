package grpc

import (
	"project_work/internal/log"
	"project_work/internal/services/grpc/monitor_v1"
	"project_work/internal/services/sysInfo"
	"time"
)

type ServerApi struct {
	monitor_v1.UnimplementedMonitorServer
	SysInfo *sysInfo.SysInfoSrv
}

func NewServerApi(sysInfo *sysInfo.SysInfoSrv) *ServerApi {
	return &ServerApi{SysInfo: sysInfo}
}

func (s *ServerApi) Connect(in *monitor_v1.RequestConnect, srv monitor_v1.Monitor_ConnectServer) error {
	for {
		data, err := s.SysInfo.ResultWork.GetAvgLastData(in.ForTheMSec)
		if err != nil {
			return err
		}
		if err := srv.Send(data); err != nil {
			log.Logger.Log.Error("Ошибка при отправки сообщения", err)
			return err
		}
		time.Sleep(time.Duration(in.EachNSec) * time.Second)
	}
}
