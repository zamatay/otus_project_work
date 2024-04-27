package models

/*
- Средняя загрузка системы (load average).

- Средняя загрузка CPU (%user_mode, %system_mode, %idle).

- Загрузка дисков:
    - tps (transfers per second);
    - KB/s (kilobytes (read+write) per second);

- Информация о дисках по каждой файловой системе:
    - использовано мегабайт, % от доступного количества;
    - использовано inode, % от доступного количества.

- Top talkers по сети:
    - по протоколам: protocol (TCP, UDP, ICMP, etc), bytes, % от sum(bytes) за последние **M**), сортируем по убыванию процента;
    - по трафику: source ip:port, destination ip:port, protocol, bytes per second (bps), сортируем по убыванию bps.

- Статистика по сетевым соединениям:
    - слушающие TCP & UDP сокеты: command, pid, user, protocol, port;
    - количество TCP соединений, находящихся в разных состояниях (ESTAB, FIN_WAIT, SYN_RCV и пр.).
*/

type SysMonitor struct {
	Time   int64
	Memory Memory
	// средняя загрузка системы
	AvgSysLoad LoadAVG // load average
	// средняя загрузка cpu
	AvgCpuLoad CpuStat // user; system; idle
	//DiskUsed   map[string]DiskInfo
	// Информация о дисках по каждой файловой системе
	DiskUsedFS map[string]DiskInfoFS //used, available, usedPercent
	DiskUsedN  map[string]DiskInfoN  //uses, free, usePercent
	// Загрузка дисков
	DiskInfo []DiskInfoAny //	Tps, KBReadS, KBWrtnS
	AvgCpu   AvgCPU
	// Статистика по сетевым соединениям
	NetInfo  []NetInfo
	StatInfo *IOStat
}
