package util

import (
	b64 "encoding/base64"
)

func ToBase64(plainText string) string {
	return b64.StdEncoding.EncodeToString([]byte(plainText))
}
