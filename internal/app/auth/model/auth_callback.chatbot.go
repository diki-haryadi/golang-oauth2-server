package model

import (
	"fmt"

	tm "golang-standards-project-layout/internal/app/token/model"
	um "golang-standards-project-layout/internal/app/user/model"
)

func BuildHandleLinkageCallbackReply(u *um.UserNoSqlSchema, t *tm.TokenNoSqlSchema) string {
	return fmt.Sprintf("Connected with username: %s and token: %s", u.Name, t.AccessToken)
}
