package usersJob

import (
	"context"

	"golang-oauth2-server/internal/pkg/wrapper"
)

func (j *job) logArticleWorker() wrapper.HandlerFunc {
	return func(ctx context.Context, args ...interface{}) (interface{}, error) {
		j.logger.Info("article log job")
		return nil, nil
	}
}
