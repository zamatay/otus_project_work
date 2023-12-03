package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	baseLog "log"
	"os"
	"project_work/internal/log"
	"time"
)

type Config struct {
	DevEnv   string         `yaml:"env" env-default:"local"`
	Grpc     GRPCConfig     `yaml:"grpc"`
	TokenTTL time.Duration  `yaml:"token_ttl"`
	Monitor  MonitorEnabled `yaml:"monitor"`
}

type MonitorEnabled struct {
	Memory      bool `yaml:"memory"`
	LoadAvg     bool `yaml:"loadAvg"`
	ProcessStat bool `yaml:"processStat"`
	Disk        bool `yaml:"disk"`
	Iostat      bool `yaml:"iostat"`
	Net         bool `yaml:"net"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func Load() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		configPath = "./config/config_local.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		baseLog.Fatal(err)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Logger.Log.Error("Config file is empty: %s", err.Error())
		baseLog.Fatal(err)
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
