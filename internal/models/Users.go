package models

import (
	"time"
)

type User struct {
	ID                     int        `json:"id" gorm:"primaryKey"`
	Name                   string     `json:"name"`
	Username               string     `json:"username"`
	Password               string     `json:"password"`
	TwoFactorSecret        string     `json:"two_factor_secret"`
	TwoFactorRecoveryCodes string     `json:"two_factor_recovery_codes"`
	TwoFactorConfirmedAt   time.Time  `json:"two_factor_confirmed_at"`
	Wallet                 float64    `json:"wallet"`
	Status                 StatusType `json:"status"`
	ProfilePhotoPath       string     `json:"profile_photo_path"`
	CreatedAt              string     `json:"created_at"`
	UpdatedAt              string     `json:"updated_at"`
	DeletedAt              string     `json:"deleted_at"`
	UserBanks              []UserBank `gorm:"foreignKey:UserID"` // One-to-many relationship with UserBank

}
