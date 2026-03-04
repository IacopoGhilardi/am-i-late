package googleMaps

import (
	"googlemaps.github.io/maps"
)

func NewClient(apiKey string) (*maps.Client, error) {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return client, nil
}
