package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	log2 "log"
	"os"
	"project_work/internal/log"
	"time"
)

type Config struct {
	Env      string        `yaml:"env" env-default:"local"`
	Grpc     GRPCConfig    `yaml:"grpc"`
	TokenTTL time.Duration `yaml:"token_ttl"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func Load() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		configPath = "./server/config/config_local.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Logger.Log.Error("config file does not exists: %s", configPath)
		log.Logger.Log.Fatal(err)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Logger.Log.Error("Config file is empty: %s", err.Error())
		log2.Fatal(err)
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
