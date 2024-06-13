package models

type AdminPermission struct {
	ID           int        `json:"id" gorm:"primaryKey"`
	AdminID      int        `json:"admin_id"`
	PermissionID int        `json:"permission_id"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
	DeletedAt    string     `json:"deleted_at"`
	Admin        Admin      `gorm:"foreignKey:AdminID;references:ID"`
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID"`
}
