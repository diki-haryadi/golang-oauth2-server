package usersKafkaConsumer

import (
	"context"
	"encoding/json"

	articleDto "golang-oauth2-server/internal/modules/article/dto"
	"golang-oauth2-server/internal/pkg/logger"
	"golang-oauth2-server/internal/pkg/wrapper"
)

func (c *consumer) createEventWorker(
	workerChan chan bool,
) wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
		defer func() {
			workerChan <- true
		}()
		for {
			msg, err := c.createEventReader.Client.FetchMessage(ctx)
			if err != nil {
				return nil, err
			}

			logger.Zap.Sugar().Infof(
				"Kafka Worker recieved message at topic/partition/offset %v/%v/%v: %s = %s\n",
				msg.Topic,
				msg.Partition,
				msg.Offset,
				string(msg.Key),
				string(msg.Value),
			)

			aDto := new(articleDto.CreateArticleRequestDto)
			if err := json.Unmarshal(msg.Value, &aDto); err != nil {
				continue
			}

			if err := c.createEventReader.Client.CommitMessages(ctx, msg); err != nil {
				return nil, err
			}
		}
	}
}
