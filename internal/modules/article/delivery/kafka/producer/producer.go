package articleKafkaProducer

import (
	"context"
	"github.com/segmentio/kafka-go"

	articleDomain "golang-oauth2-server/internal/modules/article/domain"
	kafkaProducer "golang-oauth2-server/pkg/kafka/producer"
)

type producer struct {
	createWriter *kafkaProducer.Writer
}

func NewProducer(w *kafkaProducer.Writer) articleDomain.KafkaProducer {
	return &producer{createWriter: w}
}

func (p *producer) PublishCreateEvent(ctx context.Context, messages ...kafka.Message) error {
	return p.createWriter.Client.WriteMessages(ctx, messages...)
}
