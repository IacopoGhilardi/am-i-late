package mapper

import (
	"github.com/iacopoGhilardi/amILate/internal/dto"
	"googlemaps.github.io/maps"
)

func MapToDirectionsRequest(dto dto.RouteRequestDto) *maps.DirectionsRequest {
	return &maps.DirectionsRequest{
		Origin:        dto.Origin.String(),
		Destination:   dto.Destination.String(),
		Mode:          dto.Mode,
		DepartureTime: dto.DepartureTime,
	}
}

func MapToRouteResult(route maps.Route) *dto.RouteResultDto {
	if len(route.Legs) == 0 {
		return nil
	}

	leg := route.Legs[0]

	// DurationInTraffic è 0 se non è stato passato DepartureTime
	// ma passando "now" sarà sempre popolato
	durationInTraffic := leg.DurationInTraffic
	if durationInTraffic == 0 {
		durationInTraffic = leg.Duration // fallback
	}

	return &dto.RouteResultDto{
		DurationInTraffic: durationInTraffic,
		Duration:          leg.Duration,
		DistanceMeters:    leg.Distance.Meters,
		StartAddress:      leg.StartAddress,
		EndAddress:        leg.EndAddress,
		Summary:           route.Summary,
	}
}
