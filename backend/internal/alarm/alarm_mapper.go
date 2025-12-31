package alarm

import "github.com/google/uuid"

func MapFromCreateReq(dto CreateAlarmRequestDto) *Alarm {
	return &Alarm{
		Label: dto.Label,
		Time:  dto.Time,
		//UserID:        dto.UserID,
		DestinationID: dto.DestinationID,
		Active:        true,
		PublicID:      uuid.New(),
	}
}

func MapAlarmToDto(alarm Alarm) *AlarmDto {
	return &AlarmDto{
		Label:  alarm.Label,
		Time:   alarm.Time,
		ID:     alarm.PublicID,
		Active: alarm.Active,
	}
}
