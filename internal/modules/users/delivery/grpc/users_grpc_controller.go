package usersGrpcController

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	usersV1 "golang-oauth2-server/api/users/v1"

	usersDomain "golang-oauth2-server/internal/modules/users/domain"
	usersDto "golang-oauth2-server/internal/modules/users/dto"
	usersException "golang-oauth2-server/internal/modules/users/exception"
)

type controller struct {
	useCase usersDomain.UseCase
}

func (c *controller) mustEmbedUnimplementedUsersServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewController(uc usersDomain.UseCase) usersDomain.GrpcController {
	return &controller{
		useCase: uc,
	}
}

func (c *controller) CreateUsers(ctx context.Context, req *usersV1.CreateUsersRequest) (*usersV1.CreateUsersResponse, error) {
	aDto := &usersDto.CreateUsersRequestDto{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	err := aDto.ValidateCreateArticleDto()
	if err != nil {
		return nil, usersException.CreateUsersValidationExc(err)
	}

	user, err := c.useCase.CreateUsers(ctx, aDto)
	if err != nil {
		return nil, err
	}

	return &usersV1.CreateUsersResponse{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (c *controller) GetUsersById(ctx context.Context, req *usersV1.GetUsersByIdRequest) (*usersV1.GetUsersByIdResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
