package models

import "time"

type Deposit struct {
	ID               int            `json:"id" gorm:"primaryKey"`
	UserID           int            `json:"user_id"`
	PaymentGatewayID int            `json:"payment_gateway_id"`
	Amount           float64        `json:"amount"`
	DateTime         time.Time      `json:"datetime"`
	Type             string         `json:"type"`
	Status           BankStatus     `json:"status"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
	DeletedAt        string         `json:"deleted_at"`
	PaymentGateway   PaymentGateway `gorm:"foreignKey:PaymentGatewayID;references:ID"`
}

type BankStatus string

const (
	DepositActive   BankStatus = "active"
	DepositInactive BankStatus = "inactive"
	DepositCancel   BankStatus = "cancel"
)
