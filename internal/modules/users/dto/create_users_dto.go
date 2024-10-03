package articleDto

import (
	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type CreateUsersRequestDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (caDto *CreateUsersRequestDto) ValidateCreateArticleDto() error {
	return validator.ValidateStruct(caDto,
		validator.Field(
			&caDto.Username,
			validator.Required,
			validator.Length(3, 50),
		),
		validator.Field(
			&caDto.Email,
			validator.Required,
			validator.Length(5, 100),
		),
		validator.Field(
			&caDto.Password,
			validator.Required,
			validator.Length(5, 100),
		),
	)
}

type CreateUsersResponseDto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
