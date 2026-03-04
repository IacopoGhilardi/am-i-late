package _interface

import (
	"context"

	"googlemaps.github.io/maps"
)

type MapsServiceInterface interface {
	Geocode(ctx context.Context, address string) ([]maps.GeocodingResult, error)
	Route(ctx context.Context, origin, destination maps.LatLng, mode maps.Mode) (*maps.Route, error)
}
