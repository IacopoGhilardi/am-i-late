package dto

import "github.com/google/uuid"

type DestinationDto struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name" validate:"required"`
	FormattedAddress string    `json:"formatted_address" validate:"required"`
	GooglePlaceID    string    `json:"google_place_id"`
	Latitude         float64   `json:"latitude" validate:"required"`
	Longitude        float64   `json:"longitude" validate:"required"`
}

type CreateDestinationRequestDto struct {
	Name             string  `json:"name" validate:"required"`
	FormattedAddress string  `json:"formatted_address" validate:"required"`
	Latitude         float64 `json:"latitude" validate:"required"`
	Longitude        float64 `json:"longitude" validate:"required"`
	GooglePlaceID    string  `json:"google_place_id,omitempty"`
	TransportMode    string  `json:"transport_mode" validate:"required,oneof=driving walking transit bicycling"`
}
