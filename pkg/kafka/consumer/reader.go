package kafkaConsumer

import (
	"github.com/segmentio/kafka-go"

	"golang-oauth2-server/internal/pkg/logger"
)

type Reader struct {
	Client *kafka.Reader
}

type ReaderConfig struct {
	Brokers []string
	GroupID string
	Topic   string
}

func NewKafkaReader(cfg *ReaderConfig) *Reader {
	kafkaReaderConfig := kafka.ReaderConfig{
		Brokers:     cfg.Brokers,
		GroupID:     cfg.GroupID,
		Topic:       cfg.Topic,
		Logger:      kafka.LoggerFunc(logger.Zap.Sugar().Infof),
		ErrorLogger: kafka.LoggerFunc(logger.Zap.Sugar().Errorf),
	}

	return &Reader{
		Client: kafka.NewReader(kafkaReaderConfig),
	}
}
