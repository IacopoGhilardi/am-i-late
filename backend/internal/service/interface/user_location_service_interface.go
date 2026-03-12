package _interface

import (
	"github.com/google/uuid"
)

type UserLocationServiceInterface interface {
	UpdateLocation(publicID uuid.UUID, lat, lng float64) error
}
