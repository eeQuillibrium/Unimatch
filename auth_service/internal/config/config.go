package config

import (
	"errors"
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GRPC GRPC `yaml:"grpc"`
}
type GRPC struct {
	Serverport int `yaml:"serverport"`
}

func InitConfig() (*Config, error) {
	path := fetchConfigPath()

	if path == "" {
		return nil, errors.New("empty config path")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	data, _ := os.ReadFile(path)

	var cfg Config

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "configpath", "", "path to config")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}
	return path
}
