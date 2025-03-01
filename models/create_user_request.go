package models

type CreateUserRequest struct {
	ReferrerAddress *string `json:"referrerAddress" validate:"required"`
	ReferredAddress *string `json:"referredAddress"`
}
