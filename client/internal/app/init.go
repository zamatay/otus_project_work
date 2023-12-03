package app

import (
	pb "client_monitor/internal/app/monitor_v1"
	"client_monitor/internal/config"
	"client_monitor/internal/log"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"io"
	"os"
	"strconv"
	"text/tabwriter"
)

func connectClient(cfg *config.Config, done chan struct{}) error {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	args := os.Args

	eachNSec := 15
	forTheMSec := 5
	var err error
	if len(args) >= 3 {
		if eachNSec, err = strconv.Atoi(args[1]); err != nil {
			eachNSec = 15
		}
		if forTheMSec, err = strconv.Atoi(args[2]); err != nil {
			forTheMSec = 5
		}
	}

	target := fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port)

	conn, err := grpc.Dial(target, opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	client := pb.NewMonitorClient(conn)
	request := &pb.RequestConnect{
		EachNSec:   int32(eachNSec),
		ForTheMSec: int32(forTheMSec),
	}
	response, err := client.Connect(context.Background(), request)

	go func() {
		for {
			response, err := response.Recv()
			if err == io.EOF {
				done <- struct{}{}
				return
			} else if err != nil {
				log.Logger.Fatal(err)
			}
			w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
			fmt.Fprintln(w, "Средняя загрузка системы:")
			fmt.Fprintf(w, "\t%v\n", response.SysLoadAverage)
			fmt.Fprintln(w, "Загрузка процессора:")
			fmt.Fprintf(w, "\t%d\t%d\t%d\n", response.CpuAverage.User, response.CpuAverage.System, response.CpuAverage.Idle)
			fmt.Fprintln(w, "Диск tps:")
			for _, value := range response.DiskTps {
				fmt.Fprintf(w, "\t|%v\t|%v\t|%v\n", value.Tps, value.KBReadS, value.KBWrtnS)
			}
			fmt.Fprintln(w, "Информация по дискам:")
			for _, value := range response.DiskRwPs {
				fmt.Fprintf(w, "\t|%s\t|%s\t|%s\n", value.Used, value.Available, value.UsedPercent)
			}
			fmt.Fprintln(w, "Информация по нодам:")
			for _, value := range response.DiskInfoN {
				fmt.Fprintf(w, "\t|%s\t|%s\t|%s\n", value.Uses, value.Free, value.UsePercent)
			}
			fmt.Fprintln(w, "Информация по соединениям:")
			for _, value := range response.Net.NetInfos {
				fmt.Fprintf(w, "\t|%s\t|%d\t|%d\t|%s\t|%s\n", value.State, value.RecvQ, value.SendQ, value.LocalAddress, value.PeerAddress)
			}
			fmt.Fprintln(w, "Информация по сети:")
			for _, value := range response.Net.StateInfos {
				fmt.Fprintf(w, "\t|%s\t|%d\n", value.State, value.Count)
			}
			w.Flush()
		}
	}()

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	return nil
}

func ClientRun(ctx context.Context, cancelFunc context.CancelFunc, cfg *config.Config, done chan struct{}) error {
	if err := connectClient(cfg, done); err != nil {
		log.Logger.Fatal(err)
	}
	return nil
}
