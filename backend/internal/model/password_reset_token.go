package model

import "time"

type PasswordResetToken struct {
	BaseModel
	UserID    uint       `gorm:"column:user_id;not null;index"`
	Token     string     `gorm:"column:token;not null;uniqueIndex"`
	ExpiresAt time.Time  `gorm:"column:expires_at;not null"`
	UsedAt    *time.Time `gorm:"column:used_at"`
}

func (t *PasswordResetToken) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

func (t *PasswordResetToken) IsUsed() bool {
	return t.UsedAt != nil
}
