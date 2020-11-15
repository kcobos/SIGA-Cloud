package controller

import (
	"github.com/kcobos/SIGA-Cloud/internal/errors"
	. "github.com/kcobos/SIGA-Cloud/internal/models"
)

// Places represents a list of Places.
type Places struct {
	placeList map[int]*Place
	lastID    int
}

// NewPlaces initialize Places struct
// It sets the map and the last ID
func NewPlaces() *Places {
	p := new(Places)
	p.placeList = make(map[int]*Place)
	p.lastID = -1
	return p
}

// Len returns length of place list if it is initialized
func (p *Places) Len() (int, error) {
	if p.placeList == nil {
		return 0, &errors.NotInitialized{}
	}
	return len(p.placeList), nil
}

// NewPlace appends a new place to the list of places
// Return the new id
func (p *Places) NewPlace(latitude, longitude float64, address string, parkings []*Parking) (int, error) {
	if p.placeList == nil {
		return -1, &errors.NotInitialized{}
	}

	for i := 0; i <= p.lastID; i++ {
		park := p.placeList[i]
		lat, lon := park.Location()
		if park.Address() == address ||
			(lat == latitude && lon == longitude) {
			return -1, &errors.PlaceAlreadyExists{}
		}
	}

	place, err := NewPlace(p.lastID+1, latitude, longitude, address, parkings)
	if err != nil {
		return -1, err
	}

	p.lastID++
	p.placeList[p.lastID] = place
	return p.lastID, nil
}
