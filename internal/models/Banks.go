package models

type Bank struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	NameTH    string `json:"name_th"`
	NameEN    string `json:"name_en"`
	BankShort string `json:"bank_short"`
	APIID     string `json:"api_id"`
	BankCode  string `json:"bank_code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	// DeletedAt time.Time `json:"deleted_at" gorm:"index"`
}
