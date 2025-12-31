package alarm

import "github.com/google/uuid"

type AlarmDto struct {
	ID     uuid.UUID `json:"id"`
	Label  string    `json:"label"`
	Time   string    `json:"time"`
	Active bool      `json:"is_active"`
}

type CreateAlarmRequestDto struct {
	Label         string `json:"label"`
	Time          string `json:"time"`
	UserID        uint   `json:"user_id"`
	DestinationID *uint  `json:"destination_id"`
	Active        bool   `json:"is_active"`
}
