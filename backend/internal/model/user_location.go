package model

import "time"

type UserLocation struct {
	BaseModel
	UserID             uint      `gorm:"column:user_id;not null;uniqueIndex"`
	User               User      `gorm:"foreignKey:UserID"`
	Latitude           float64   `gorm:"column:latitude;not null"`
	Longitude          float64   `gorm:"column:longitude;not null"`
	LastNotificationAt time.Time `gorm:"column:last_notification_at;not null"`
}
