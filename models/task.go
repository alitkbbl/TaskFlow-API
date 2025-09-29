package models

type Task struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Status      string `gorm:"default:todo" json:"status"`
	UserID      string `gorm:"not null" json:"user_id"`
	CreatedAt   int64  `gorm:"autoCreateTime" json:"created_at"`
}
