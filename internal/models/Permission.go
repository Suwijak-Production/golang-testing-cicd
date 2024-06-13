package models

type Permission struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	NameTH    string `json:"name_th"`
	NameEN    string `json:"name_en"`
	Sort      int    `json:"sort"`
	Code      string `json:"code"`
	Group     string `json:"group"`
	Type      string `json:"type"`
	TypeTH    string `json:"type_th"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
