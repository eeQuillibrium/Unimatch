package app

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	kafkaClient "github.com/eeQuillibrium/Unimatch/pkg/kafka"
	"github.com/eeQuillibrium/Unimatch/pkg/logger"
	grpcapp "github.com/eeQuillibrium/Unimatch/profile_service/internal/app/grpc"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/config"
	kafkaReader "github.com/eeQuillibrium/Unimatch/profile_service/internal/delivery/kafka"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/repository"
	"github.com/eeQuillibrium/Unimatch/profile_service/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

type App struct {
	log       *logger.Logger
	cfg       *config.Config
	cg        kafkaClient.ConsumerGroup
	kafkaConn *kafka.Conn
}

func NewApp(log *logger.Logger, cfg *config.Config) *App {
	return &App{log: log, cfg: cfg}
}

func (a *App) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.cfg.PostgresDB.Host, a.cfg.PostgresDB.Port, a.cfg.PostgresDB.Username, a.cfg.PostgresDB.DBName, a.cfg.PostgresDB.Password, a.cfg.PostgresDB.SSLMode)

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		return err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", a.cfg.RedisDB.Host, a.cfg.RedisDB.Port),
		Password: a.cfg.RedisDB.Password,
		DB:       a.cfg.RedisDB.DB,
	})
	repo := repository.NewRepository(a.log, a.cfg, db, rdb)

	services := service.NewService(a.log, repo)

	grpcApp := grpcapp.NewGRPCApp(a.log, a.cfg, services.ProfileProvider)

	if err := grpcApp.Run(); err != nil {
		a.log.Warnf("gRPCApp.Run(): %v", err)
	}
	msgReader := kafkaReader.NewMessageReader(a.log, &a.cfg.KafkaTopics, services)

	a.cg = kafkaClient.NewConsumerGroup(a.log, a.cfg.Kafka.Brokers, a.cfg.Kafka.GroupID)

	if err := a.setKafkaConn(ctx, &a.cfg.Kafka); err != nil {
		a.log.Warnf("kafka.NewKafkaConn(): %w", err)
	}
	defer a.kafkaConn.Close()

	go a.cg.ConsumeTopics(ctx, a.getKafkaTopics(), kafkaReader.PoolSize, msgReader.MessageReaderWorker)

	<-ctx.Done()

	return nil
}
