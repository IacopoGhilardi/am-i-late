package service

import (
	"context"
	"fmt"

	"github.com/iacopoGhilardi/amILate/internal/dto"
	"github.com/iacopoGhilardi/amILate/internal/mapper"
	"googlemaps.github.io/maps"
)

type MapService struct {
	client *maps.Client
}

func NewMapService(client *maps.Client) *MapService {
	return &MapService{client: client}
}

func (s *MapService) Geocode(ctx context.Context, address string) ([]maps.GeocodingResult, error) {
	return s.client.Geocode(ctx, &maps.GeocodingRequest{Address: address})
}

func (s *MapService) Route(ctx context.Context, req dto.RouteRequestDto) (*dto.RouteResultDto, error) {
	resp, _, err := s.client.Directions(ctx, mapper.MapToDirectionsRequest(req))
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, fmt.Errorf("nessun percorso trovato")
	}
	return mapper.MapToRouteResult(resp[0]), nil
}
