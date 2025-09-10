package user

import (
	"github.com/iacopoGhilardi/amILate/internal/commons"
)

type User struct {
	commons.BaseModel
	Email    string `gorm:"column:email;type:varchar(255);unique;not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
}
