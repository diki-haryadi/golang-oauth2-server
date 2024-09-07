package util

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"github.com/rs/zerolog/log"
	"golang-standards-project-layout/internal/constants"
	errlib "golang-standards-project-layout/pkg/err"
	"golang-standards-project-layout/pkg/response"
)

func NewRequiredFieldErr(requiredField string) errlib.Error {
	err := fmt.Errorf(constants.ErrRequiredField, requiredField)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func NewDataNotFoundErr(entityName string) errlib.Error {
	err := fmt.Errorf(constants.ErrDataNotFound, entityName)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func NewInvalidFieldFormatErr(invalidField string) errlib.Error {
	err := fmt.Errorf(constants.ErrRequiredField, invalidField)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func NewInvalidEnumErr(value interface{}, enumName string) errlib.Error {
	err := fmt.Errorf(constants.ErrInvalidEnum, value, enumName)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func NewBadRequestErr(msg string) errlib.Error {
	err := fmt.Errorf(msg)

	log.Err(err).Msg(err.Error())
	return errlib.NewError(err, errlib.ErrBadRequest)
}

func ReturnErrorToFiberResponse(fc *fiber.Ctx, err error) error {
	e := errlib.GetError(err)
	res := response.NewJSONResponse().APIStatusCreated().SetError(e)
	//.WithErrorString(e.Error()).WithStatusCode(e.HttpStatusCode())
	return fc.JSON(res)
}
