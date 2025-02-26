package models

type Referral struct {
    ID         uint  `gorm:"primarykey" json:"id"`
    CreatedAt  uint  `json:"created_at,omitempty"`
    UpdatedAt  uint  `json:"updated_at,omitempty"`
    DeletedAt  uint  `json:"deleted_at,omitempty" gorm:"index"`
    ReferrerID uint  `json:"referrer_id,omitempty"`
    RefereeID  uint  `json:"referee_id,omitempty"`
    Referrer   *User `json:"referrer,omitempty"`
    Referee    *User `json:"referee"`
}