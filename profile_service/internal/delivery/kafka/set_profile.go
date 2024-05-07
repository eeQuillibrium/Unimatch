package kafka

import (
	"context"
	"time"

	"github.com/avast/retry-go"
	kafkaMessages "github.com/eeQuillibrium/Unimatch/proto/gen/go/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (w *messageReader) SetProfileProcessor(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	w.log.Info("message delivered to setProfile()")

	var profile kafkaMessages.Profile

	if err := proto.Unmarshal(m.Value, &profile); err != nil {
		w.log.Warnf("proto.Unmarshal(): %v", err)
		r.CommitMessages(ctx, m)
		return
	}

	if err := retry.Do(func() error {
		return w.services.ProfileProvider.StoreProfile(ctx, &profile)
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		w.log.Warnf("StoreProfile(): %v", err)
		return
	}

	w.log.Infof("user with name: %s, age: %d, about: %s was stored", profile.Name, profile.Age, profile.About)

	r.CommitMessages(ctx, m)
}
