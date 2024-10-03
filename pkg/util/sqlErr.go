package util

import (
	"fmt"
	"strings"

	"golang-oauth2-server/pkg/custerr"
	"golang-oauth2-server/pkg/response"
)

var (
	invalidConstraint = "Error 1452"
	duplicateEntry    = "Error 1062"
)

type ErrSqlStatement struct {
	ErrMessage  string
	Err         error
	Constraints []string
}

func (e ErrSqlStatement) Error() error {
	if strings.HasPrefix(e.Err.Error(), invalidConstraint) {
		for _, constraint := range e.Constraints {
			if strings.ContainsAny(e.Err.Error(), constraint) {
				return &custerr.ErrChain{
					Message: e.ErrMessage,
					Type:    response.ErrBadRequest,
					Cause:   fmt.Errorf("invalid %s", constraint),
				}
			}
		}
	}
	if strings.HasPrefix(e.Err.Error(), duplicateEntry) {
		return &custerr.ErrChain{
			Message: e.ErrMessage,
			Type:    response.ErrBadRequest,
			Cause:   fmt.Errorf("id already exist"),
		}
	}
	return &custerr.ErrChain{
		Message: e.ErrMessage,
		Type:    response.ErrBadRequest,
		Cause:   e.Err,
	}
}
