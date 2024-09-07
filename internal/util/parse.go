package util

import (
	"strconv"
	"time"
)

func ParseStringToInt64(value string, fieldName string) (int64, error) {
	parsedVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, NewInvalidFieldFormatErr(fieldName)
	}

	return parsedVal, err
}

func ParseStringToTime(value string, timeFormat string, fieldName string) (time.Time, error) {
	parsedVal, err := time.Parse(timeFormat, value)
	if err != nil {
		return time.Time{}, NewInvalidFieldFormatErr(fieldName)
	}

	return parsedVal, nil
}
