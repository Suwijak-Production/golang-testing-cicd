package models

import "time"

type Transaction struct {
	ID               int               `json:"id" gorm:"primaryKey"`
	UserID           int               `json:"user_id"`
	AdminID          int               `json:"admin_id"`
	PaymentGatewayID int               `json:"payment_gateway_id"`
	DepositID        int               `json:"deposit_id"`
	WithdrawID       int               `json:"withdraw_id"`
	DateTime         time.Time         `json:"datetime"`
	Amount           float64           `json:"amount"`
	Status           TransactionStatus `json:"status"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
	DeletedAt        string            `json:"deleted_at"`
	User             User              `gorm:"foreignKey:UserID;references:ID"`
	PaymentGateway   PaymentGateway    `gorm:"foreignKey:PaymentGatewayID;references:ID"`
	Admin            Admin             `gorm:"foreignKey:AdminID;references:ID"`
	Deposit          Deposit           `gorm:"foreignKey:DepositID;references:ID"`
	Withdraw         Withdraw          `gorm:"foreignKey:WithdrawID;references:ID"`
}

type TransactionStatus string

const (
	Waiting TransactionStatus = "waiting"
	Process TransactionStatus = "process"
	Success TransactionStatus = "success"
	Error   TransactionStatus = "error"
	Reject  TransactionStatus = "reject"
	Other   TransactionStatus = "other"
)
