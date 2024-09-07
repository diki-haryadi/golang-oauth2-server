package custerr

import (
	"fmt"
)

type ErrChain struct {
	Message string
	Cause   error
	Fields  map[string]interface{}
	Type    error
}

func (ec ErrChain) Error() string {
	coz := ""
	fields := ""
	if ec.Cause != nil {
		coz = fmt.Sprint(" because {", ec.Cause.Error(), "}")
		if len(ec.Fields) > 0 {
			fields = fmt.Sprintf(" with Fields {%+v}", ec.Fields)
		}
	}

	return fmt.Sprint(ec.Message, coz, fields)
}

func (ec ErrChain) SetFields(key string, val string) ErrChain {
	if ec.Fields == nil {
		ec.Fields = map[string]interface{}{}
	}

	ec.Fields[key] = val
	return ec
}

type InvalidError struct {
	message string
}

func (ie *InvalidError) Error() string {
	return ie.message
}

func NewInvalidError(msg string) *InvalidError {
	return &InvalidError{message: msg}
}

func NewInvalidErrorf(msg string, args ...interface{}) *InvalidError {
	return NewInvalidError(fmt.Sprintf(msg, args...))
}
