package models

type User struct {
	ID           string `gorm:"primaryKey" json:"id"`
	Email        string `gorm:"uniqueIndex;not null" json:"email"`
	Username     string `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string `gorm:"not null" json:"-"`
	CreatedAt    int64  `gorm:"autoCreateTime" json:"created_at"`
}
