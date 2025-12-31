package commons

import "time"

type BaseModel struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type BaseModelWithSafeDelete struct {
	BaseModel
	DeletedAt *time.Time `gorm:"column:deleted_at;index"`
}

type Location struct {
	Latitude  float64 `gorm:"column:latitude;type:decimal(9,6)"`
	Longitude float64 `gorm:"column:longitude;type:decimal(9,6)"`
}
