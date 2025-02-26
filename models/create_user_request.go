package models

type CreateUserRequest struct {
	Address string `json:"address" validate:"required"`
}
