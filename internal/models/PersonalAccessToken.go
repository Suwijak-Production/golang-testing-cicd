package models

type PersonalAccessToken struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	TokenableType string `json:"tokenable_type"`
	TokenableID   int64  `json:"tokenable_id"`
	Name          string `json:"name"`
	Token         string `json:"token"`
	Abilities     string `json:"abilities"`
	LastUsedAt    string `json:"last_used_at"`
	ExpiresAt     string `json:"expires_at"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
