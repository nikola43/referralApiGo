package services

import (
	"errors"
	"time"

	"github.com/nikola43/pdexrefapi/db"
	"github.com/nikola43/pdexrefapi/models"
)

func CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		Address:   req.Address,
		CreatedAt: uint(time.Now().Unix()),
		UpdatedAt: uint(time.Now().Unix()),
	}

	tx := db.GormDB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func GetUser(address string) (*models.User, error) {
	user := &models.User{}

	tx := db.GormDB.Where("address = ?", address).First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func AddReferral(req *models.AddReferralRequest) (*models.Referral, error) {

	// check if referrer and referee are the same
	if req.ReferrerAddress == req.RefereeAddress {
		return nil, errors.New("referrer and referee cannot be the same")
	}

	// Check if referrer exists
	referrer := &models.User{}
	tx := db.GormDB.Where("address = ?", req.ReferrerAddress).First(referrer)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// Check if referee exists
	referee := &models.User{}
	tx = db.GormDB.Where("address = ?", req.RefereeAddress).First(referee)
	if tx.Error != nil {
		return nil, tx.Error
	}

	ReferrerID := referrer.ID
	RefereeID := referee.ID

	// check if referral already exists
	referral := &models.Referral{}
	tx = db.GormDB.Where("referrer_id = ? AND referee_id = ?", ReferrerID, RefereeID).First(referral)
	if tx.Error == nil {
		return nil, errors.New("referral already exists")
	}

	referral = &models.Referral{
		ReferrerID: ReferrerID,
		RefereeID:  RefereeID,
		CreatedAt:  uint(time.Now().Unix()),
		UpdatedAt:  uint(time.Now().Unix()),
	}

	tx = db.GormDB.Create(referral)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// return referrer with all referrals
	tx = db.GormDB.Preload("Referrals").Where("id = ?", ReferrerID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return referral, nil
}

func GetUserWithReferrals(address string) (*models.User, error) {
	user := &models.User{}

	tx := db.GormDB.Preload("Referrals.Referrer").Preload("Referrals.Referee").Where("address = ?", address).First(user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	// create a new slice of referrals with only the required fields
	referrals := make([]models.Referral, len(user.Referrals))
	for i, referral := range user.Referrals {
		referrals[i] = models.Referral{
			ID:      referral.ID,
			Referee: referral.Referee,
		}
	}

	// set the new referrals slice to the user
	user.Referrals = referrals

	return user, nil
}
