package articleUseCase

import (
	"context"
	clientDomain "golang-oauth2-server/internal/modules/client/domain"
	"golang-oauth2-server/pkg/util"
)

type useCase struct {
	repository clientDomain.Repository
}

func NewUseCase(
	repository clientDomain.Repository,
) clientDomain.UseCase {
	return &useCase{
		repository: repository,
	}
}

func (uc *useCase) AuthClient(ctx context.Context, clientID, secret string) (*clientDomain.Client, error) {
	client, err := uc.repository.FindClientByClientID(ctx, clientID)
	if err != nil {
		return nil, err
	}

	if util.VerifyPassword(client.Secret, secret) != nil {
		return nil, err
	}
	return client, nil
}

func (uc *useCase) CreateClient(ctx context.Context, clientID, secret, redirectURI string) (*clientDomain.Client, error) {
	client, err := uc.repository.CreateClientCommon(ctx, clientID, secret, redirectURI)
	if err != nil {
		return nil, err
	}
	return client, err
}

func (uc *useCase) ClientExists(ctx context.Context, clientID string) bool {
	_, err := uc.repository.FindClientByClientID(ctx, clientID)
	return err == nil
}
