package models

type PaymentGateway struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Endpoint  string     `json:"endpoint"`
	Token     string     `json:"token"`
	Name      string     `json:"name"`
	Status    StatusType `json:"status"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	DeletedAt string     `json:"deleted_at"`
}

type StatusType string

const (
	Active   StatusType = "active"
	Inactive StatusType = "inactive"
)
