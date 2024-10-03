package kafkaHealthCheckUseCase

import (
	"github.com/segmentio/kafka-go"
	"golang-oauth2-server/config"

	healthCheckDomain "golang-oauth2-server/internal/modules/health_check/domain"
)

type useCase struct{}

func NewUseCase() healthCheckDomain.KafkaHealthCheckUseCase {
	return &useCase{}
}

func (uc *useCase) Check() bool {
	brokers := kafka.TCP(config.BaseConfig.Kafka.ClientBrokers...)

	conn, err := kafka.Dial(brokers.Network(), brokers.String())
	if err != nil {
		return false
	}

	_ = conn.Close()

	return true
}
