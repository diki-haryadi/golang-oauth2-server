package model

import (
	sm "golang-standards-project-layout/internal/app/session/model"
	tm "golang-standards-project-layout/internal/app/token/model"
	um "golang-standards-project-layout/internal/app/user/model"
)

type CommandMetadata struct {
	User      um.UserNoSqlSchema
	Session   sm.SessionNoSqlSchema
	HostToken tm.TokenNoSqlSchema
}
