package sampleExtServiceDomain

import (
	"context"

	articleV1 "golang-oauth2-server/api/article/v1"
)

type SampleExtServiceUseCase interface {
	CreateUsers(ctx context.Context, req *articleV1.CreateArticleRequest) (*articleV1.CreateArticleResponse, error)
}
