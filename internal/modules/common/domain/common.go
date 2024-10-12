package commonDomain

import (
	"github.com/google/uuid"
	"time"
)

// CommonModel uses uuid's for ID, generated in go
type CommonModel struct {
	ID        uuid.UUID  `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

// TimestampModel
type TimestampModel struct {
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

// EmailTokenModel ia an abstract model which can be used for objects from which
// we derive redirect emails (email confirmation, password reset and such)
type EmailTokenModel struct {
	CommonModel
	Reference   string     `db:"reference"`
	EmailSent   bool       `db:"email_sent"`
	EmailSentAt *time.Time `db:email_sent_at`
	ExpiresAt   time.Time  `db:"expires_at"`
}
