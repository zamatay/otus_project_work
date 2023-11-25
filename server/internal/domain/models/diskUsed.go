package models

type DiskInfo struct {
	ReadCount  uint64 `json:"readCount"`
	WriteCount uint64 `json:"writeCount"`
	Tps        uint64 `json:"tps"` //+
	Name       string `json:"name"`
	ReadBytes  uint64 `json:"ReadBytes"`  //+
	WriteBytes uint64 `json:"WriteBytes"` //+
	ReadTime   uint64 `json:"ReadTime"`
	WriteTime  uint64 `json:"WriteTime"`
	IoTime     uint64 `json:"IoTime"`
}

type DiskInfoFS struct {
	Fs          string
	Kblock      string
	Used        string //+
	Available   string //+
	UsedPercent string //+
}

type DiskInfoN struct {
	Fs         string
	Node       int64
	Uses       int64
	Free       int64
	UsePercent string
	Mount      string
}
