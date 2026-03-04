package _interface

import (
	"context"

	"github.com/iacopoGhilardi/amILate/internal/dto"
	"googlemaps.github.io/maps"
)

type MapsServiceInterface interface {
	Geocode(ctx context.Context, address string) ([]maps.GeocodingResult, error)
	Route(ctx context.Context, req dto.RouteRequestDto) (*dto.RouteResultDto, error)
}
