package config

import (
	"errors"
	"flag"
	"os"

	"github.com/eeQuillibrium/Unimatch/pkg/kafka"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Kafka       kafka.Config `yaml:"kafka"`
	KafkaTopics KafkaTopics  `yaml:"kafkaTopics"`
	PostgresDB  PostgresDB   `yaml:"postgresDB"`
}

type KafkaTopics struct {
	SetProfile kafka.TopicConfig `yaml:"setProfile"`
}
type PostgresDB struct {
	Host       string   `yaml:"host"`
	Password   string   `yaml:"password"`
	Port       int      `yaml:"port"`
	SSLMode    string   `yaml:"sslmode"`
	DBName     string   `yaml:"dbname"`
	Username   string   `yaml:"username"`
	TableNames []string `yaml:"tableNames"`
}

func InitConfig() (*Config, error) {
	path := fetchConfigPath()

	if path == "" {
		return nil, errors.New("empty config path")
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "cfgpath", "config/config.yaml", "specify config path")

	if path == "" {
		path = os.Getenv("CONFIG_PROFILE_PATH")
	}

	return path
}
