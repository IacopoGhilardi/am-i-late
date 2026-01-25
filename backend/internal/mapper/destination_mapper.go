package mapper

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/commons"
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/model"
)

func MapDestinationToDto(dest model.Destination) *dto.DestinationDto {
	return &dto.DestinationDto{
		Name: dest.Name,
	}
}

func MapFromCreateReq(dto dto.CreateDestinationRequestDto) *model.Destination {
	return &model.Destination{
		PublicID: uuid.New(),
		Name:     dto.Name,
		AddressComponents: model.AddressComponents{
			FormattedAddress: dto.FormattedAddress,
			GooglePlaceId:    dto.GooglePlaceID,
		},
		Location: commons.Location{
			Latitude:  dto.Latitude,
			Longitude: dto.Longitude,
		},
		TimeZone: commons.DefaultTimeZone,
	}
}
