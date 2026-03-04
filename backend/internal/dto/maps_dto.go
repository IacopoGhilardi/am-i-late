package dto

import (
	"time"

	"googlemaps.github.io/maps"
)

type RouteRequestDto struct {
	Origin        maps.LatLng
	Destination   maps.LatLng
	Mode          maps.Mode
	DepartureTime string // "now" o unix timestamp
}

type RouteResultDto struct {
	// Durata con traffico in tempo reale — campo principale
	DurationInTraffic time.Duration
	// Durata senza traffico — fallback
	Duration time.Duration
	// Distanza in metri
	DistanceMeters int
	// Indirizzo di partenza e arrivo
	StartAddress string
	EndAddress   string
	// Strada principale del percorso
	Summary string
}
