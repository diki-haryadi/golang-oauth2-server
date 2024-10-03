package articleException

import (
	errorList "golang-oauth2-server/pkg/constant/error/error_list"
	customErrors "golang-oauth2-server/pkg/error/custom_error"
	errorUtils "golang-oauth2-server/pkg/error/error_utils"
)

func CreateUsersValidationExc(err error) error {
	ve, ie := errorUtils.ValidationErrorHandler(err)
	if ie != nil {
		return ie
	}

	validationError := errorList.InternalErrorList.ValidationError
	return customErrors.NewValidationError(validationError.Msg, validationError.Code, ve)
}

func UsersBindingExc() error {
	usersBindingError := errorList.InternalErrorList.UsersExceptions.BindingError
	return customErrors.NewBadRequestError(usersBindingError.Msg, usersBindingError.Code, nil)
}
