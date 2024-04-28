package config

import (
	"errors"
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GRPC       GRPC       `yaml:"grpc"`
	PostgresDB PostgresDB `yaml:"postgresdb"`
}
type GRPC struct {
	Serverport int `yaml:"serverport"`
}
type PostgresDB struct {
	Username string `yaml:"username"`
	DBName   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
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

	flag.StringVar(&path, "cfgpath", "", "path to config")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_AUTH_PATH")
	}
	return path
}
