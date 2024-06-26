package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	DevEnv   string        `yaml:"env" env-default:"local"`
	Grpc     GRPCConfig    `yaml:"grpc"`
	TokenTTL time.Duration `yaml:"token_ttl"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout"`
}

func Load() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		configPath = "./config/config_local.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exists: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Panicf("Config file is empty: %s", err.Error())
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
