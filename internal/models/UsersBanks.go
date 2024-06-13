package models

type UserBank struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	UserID      int        `json:"user_id"`
	BankID      int        `json:"bank_id"`
	Account     string     `json:"account"`
	AccountName string     `json:"account_name"`
	Nickname    string     `json:"nickname"`
	Status      BankStatus `json:"status"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
	DeletedAt   string     `json:"deleted_at"`
	Bank        Bank       `gorm:"foreignKey:BankID;references:ID"`
}
