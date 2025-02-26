package models

type User struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	Address   string     `gorm:"unique" json:"address"`
	CreatedAt uint     `json:"created_at"`
	UpdatedAt uint     `json:"updated_at"`
	DeletedAt uint     `json:"deleted_at" gorm:"index"`
	Referrals []Referral `json:"referrals" gorm:"foreignKey:ReferrerID"` // A user can refer many other users
}
