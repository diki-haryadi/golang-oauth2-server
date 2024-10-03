package articleHttpController

import (
	"github.com/labstack/echo/v4"
	"net/http"

	articleDomain "golang-oauth2-server/internal/modules/article/domain"
	articleDto "golang-oauth2-server/internal/modules/article/dto"
	articleException "golang-oauth2-server/internal/modules/article/exception"
)

type controller struct {
	useCase articleDomain.UseCase
}

func NewController(uc articleDomain.UseCase) articleDomain.HttpController {
	return &controller{
		useCase: uc,
	}
}

func (c controller) CreateUsers(ctx echo.Context) error {
	aDto := new(articleDto.CreateArticleRequestDto)
	if err := ctx.Bind(aDto); err != nil {
		return articleException.ArticleBindingExc()
	}

	if err := aDto.ValidateCreateArticleDto(); err != nil {
		return articleException.CreateArticleValidationExc(err)
	}

	article, err := c.useCase.CreateUsers(ctx.Request().Context(), aDto)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, article)
}
