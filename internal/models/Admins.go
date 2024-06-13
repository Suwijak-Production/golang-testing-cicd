package models

import "time"

type Admin struct {
	ID              int        `json:"id" gorm:"primaryKey"`
	Username        string     `json:"username"`
	Password        string     `json:"password"`
	Name            string     `json:"name"`
	Phone           string     `json:"phone"`
	Twofactor       string     `json:"twofactor"`
	TwoFactorStatus StatusType `json:"twofactor_status" gorm:"column:twofactor_status"`
	Status          StatusType `json:"status"`
	CreatedAt       time.Time  `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       string     `json:"deleted_at" gorm:"default:null"`
}
