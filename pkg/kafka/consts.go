package kafka

import "time"

const (
	minBytes               = 10e3 // 10KB
	maxBytes               = 10e6 // 10MB
	queueCapacity          = 100
	heartbeatInterval      = 3 * time.Second
	commitInterval         = 0
	partitionWatchInterval = 5 * time.Second
	maxAttempts            = 3
	dialTimeout            = 3 * time.Minute
	maxWait                = 1 * time.Second

	producerReadTimeout  = 10 * time.Second
	producerWriteTimeout = 10 * time.Second
	producerRequiredAcks = -1
	producerMaxAttempts  = 3
)
