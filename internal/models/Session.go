package models

type Session struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	UserID       int64  `json:"user_id"`
	IPAddress    string `json:"ip_address"`
	UserAgent    string `json:"user_agent"`
	Payload      string `json:"payload"`
	LastActivity int    `json:"last_activity"`
	User         User   `gorm:"foreignKey:UserID;references:ID"`
}
