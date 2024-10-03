package usersHttpController

import (
	"net/http"

	"github.com/labstack/echo/v4"

	usersDomain "golang-oauth2-server/internal/modules/users/domain"
	usersDto "golang-oauth2-server/internal/modules/users/dto"
	usersException "golang-oauth2-server/internal/modules/users/exception"
)

type controller struct {
	useCase usersDomain.UseCase
}

func NewController(uc usersDomain.UseCase) usersDomain.HttpController {
	return &controller{
		useCase: uc,
	}
}

func (c controller) CreateUsers(ctx echo.Context) error {
	aDto := new(usersDto.CreateUsersRequestDto)
	if err := ctx.Bind(aDto); err != nil {
		return usersException.UsersBindingExc()
	}

	if err := aDto.ValidateCreateArticleDto(); err != nil {
		return usersException.CreateUsersValidationExc(err)
	}

	article, err := c.useCase.CreateUsers(ctx.Request().Context(), aDto)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, article)
}
