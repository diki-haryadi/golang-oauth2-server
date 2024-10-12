package tokenDomain

import (
	"database/sql"
	clientDomain "golang-oauth2-server/internal/modules/client/domain"
	"golang-oauth2-server/internal/modules/common/domain"
	usersDomain "golang-oauth2-server/internal/modules/users/domain"
	"time"
)

type RefreshToken struct {
	commonDomain.CommonModel
	ClientID  sql.NullString `db:"client_id"`
	UserID    sql.NullString `db:"user_id"`
	Client    *clientDomain.Client
	User      *usersDomain.Users
	Token     string    `sql:"token"`
	ExpiresAt time.Time `sql:"expires_at"`
	Scope     string    `sql:"scope"`
}
