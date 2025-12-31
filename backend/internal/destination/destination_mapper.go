package destination

import (
	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/commons"
)

func MapDestinationToDto(dest Destination) *DestinationDto {
	return &DestinationDto{
		Name: dest.Name,
	}
}

func MapFromCreateReq(dto CreateDestinationRequestDto) *Destination {
	return &Destination{
		PublicID: uuid.New(),
		Name:     dto.Name,
		AddressComponents: AddressComponents{
			FormattedAddress: dto.FormattedAddress,
			GooglePlaceId:    dto.GooglePlaceID,
		},
		Location: commons.Location{
			Latitude:  dto.Latitude,
			Longitude: dto.Longitude,
		},
		TransportMode: dto.TransportMode,
		TimeZone:      commons.DefaultTimeZone,
	}
}
