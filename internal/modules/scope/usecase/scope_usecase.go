package articleUseCase

import (
	"context"
	"github.com/pkg/errors"

	scopeDomain "golang-oauth2-server/internal/modules/scope/domain"
)

type useCase struct {
	repository scopeDomain.Repository
}

func NewUseCase(
	repository scopeDomain.Repository,
) scopeDomain.UseCase {
	return &useCase{
		repository: repository,
	}
}

var (
	// ErrInvalidScope ...
	ErrInvalidScope = errors.New("Invalid scope")
)

func (uc *useCase) GetScope(ctx context.Context, requestScope string) (string, error) {
	if requestScope == "" {
		scope, err := uc.repository.GetDefaultScope(ctx)
		if err != nil {
			return "", err
		}
		return scope, nil
	}

	if scope, err := uc.repository.ScopeExists(ctx, requestScope); scope {
		if err != nil {
			return "", err
		}
		return requestScope, nil
	}
	return "", ErrInvalidScope
}
