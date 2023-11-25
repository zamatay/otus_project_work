package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"project_work/internal/config"
	pb "project_work/internal/services/grpc/monitor_v1"
	"project_work/internal/services/sysInfo"
)

func Serve(cfg config.GRPCConfig, info *sysInfo.SysInfoSrv) {
	// Создаём новый сервер с единственным интерсептором
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	pb.RegisterMonitorServer(grpcServer, pb.NewServerApi(info))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
