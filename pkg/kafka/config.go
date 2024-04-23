package kafka

type Config struct {
	Brokers    []string `yaml:"brokers"`
	GroupID    string   `yaml:"groupID"`
	InitTopics bool     `yaml:"initTopics"`
}

type TopicConfig struct {
	TopicName         string `yaml:"topicName"`
	Partitions        int    `yaml:"partitions"`
	ReplicationFactor int    `yaml:"replicationFactor"`
}
