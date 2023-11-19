package models

/*

- Top talkers по сети:
- по протоколам: protocol (TCP, UDP, ICMP, etc), bytes, % от sum(bytes) за последние **M**), сортируем по убыванию процента;
- по трафику: source ip:port, destination ip:port, protocol, bytes per second (bps), сортируем по убыванию bps.

- Статистика по сетевым соединениям:
- слушающие TCP & UDP сокеты: command, pid, user, protocol, port;
- количество TCP соединений, находящихся в разных состояниях (ESTAB, FIN_WAIT, SYN_RCV и пр.).

*/

type SysMonitor struct {
	Time       int64
	Memory     Memory
	AvgSysLoad LoadAVG
	AvgCpuLoad CpuStat
	DiskUsed   map[string]DiskInfo
	DiskUsedFS map[string]DiskInfoFS
	DiskUsedN  map[string]DiskInfoN
	DiskInfo   []DiskInfoAny
	AvgCpu     AvgCPU
	NetInfo    []NetInfo
	StatInfo   IOStat
}
