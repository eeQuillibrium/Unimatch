package config

import (
	"errors"
	"flag"
	"os"

	"github.com/eeQuillibrium/Unimatch/pkg/kafka"
	"gopkg.in/yaml.v3"
)

type Config struct {
	GRPC        GRPC         `yaml:"grpc"`
	Http        Http         `yaml:"http"`
	Kafka       kafka.Config `yaml:"kafka"`
	KafkaTopics KafkaTopics  `yaml:"kafkatopics"`
	AssetsPath  string       `yaml:"assetsPath"`
}
type GRPC struct {
	AuthPort int `yaml:"authport"`
}
type Http struct {
	Port int `yaml:"port"`
}
type KafkaTopics struct {
	SetProfile kafka.TopicConfig `yaml:"setprofile"`
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
		path = os.Getenv("CONFIG_GATEWAY_PATH")
	}
	return path
}
