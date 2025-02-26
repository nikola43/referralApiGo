package models

type AddReferralRequest struct {
	ReferrerAddress string `json:"referrerAddress" validate:"required"`
	RefereeAddress  string `json:"refereeAddress" validate:"required"`
}
