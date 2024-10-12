package usersDomain

import (
	"github.com/google/uuid"
	"golang-oauth2-server/internal/modules/common/domain"
)

type Role struct {
	commonDomain.TimestampModel
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}
