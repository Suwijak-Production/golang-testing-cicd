package models

type FailedJob struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	UUID       string `json:"uuid"`
	Connection string `json:"connection"`
	Queue      string `json:"queue"`
	Payload    string `json:"payload"`
	Exception  string `json:"exception"`
	FailedAt   string `json:"failed_at"`
}
