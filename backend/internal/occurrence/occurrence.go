package occurrence

import (
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/destination"
	"github.com/iacopoGhilardi/amILate/internal/route"
	"github.com/iacopoGhilardi/amILate/internal/user"
)

var recurrences = []string{"daily", "weekly", "monthly"}

type Occurrence struct {
	commons.BaseModelWithSafeDelete

	User        user.User               `gorm:"foreignKey:UserID"`
	Destination destination.Destination `gorm:"foreignKey:DestinationID"`
	Route       route.Route             `gorm:"foreignKey:RouteID"`
	Recurrence  Recurrence              `gorm:"embedded"`
}

type Recurrence struct {
	Type          string   `gorm:"column:type;type:varchar(255);not null"`
	DaysOfTheWeek []string `gorm:"column:days_of_the_week;type:varchar(255);not null"`
	UntilDate     string   `gorm:"column:until_date;type:varchar(255);not null"`
}
