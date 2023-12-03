package models

type DiskInfoAny struct {
	DiskDevice string  `json:"disk_device"`
	Tps        float64 `json:"tps"`
	KBReadS    float64 `json:"kB_read/s"`
	KBWrtnS    float64 `json:"kB_wrtn/s"`
	KBDscdS    float64 `json:"kB_dscd/s"`
	KBRead     int     `json:"kB_read"`
	KBWrtn     int     `json:"kB_wrtn"`
	KBDscd     int     `json:"kB_dscd"`
}

type AvgCPU struct {
	User   float64 `json:"user"`
	Nice   float64 `json:"nice"`
	System float64 `json:"system"`
	Iowait float64 `json:"iowait"`
	Steal  float64 `json:"steal"`
	Idle   float64 `json:"idle"`
}

type IOStat struct {
	Sysstat struct {
		Hosts []struct {
			Nodename     string `json:"nodename"`
			Sysname      string `json:"sysname"`
			Release      string `json:"release"`
			Machine      string `json:"machine"`
			NumberOfCpus int    `json:"number-of-cpus"`
			Date         string `json:"date"`
			Statistics   []struct {
				AvgCPU AvgCPU        `json:"avg-cpu"`
				Disk   []DiskInfoAny `json:"disk"`
			} `json:"statistics"`
		} `json:"hosts"`
	} `json:"sysstat"`
}
