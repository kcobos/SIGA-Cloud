package model

import (
	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

// Place represents a parking location
type Place struct {
	id                  int
	latitude, longitude float64
	address             string
	parkings            []int
	freeParkings        int
}

// NewPlace create a place object
func NewPlace(id int, latitude, longitude float64, address string, parkings []*Parking) (*Place, error) {
	if len(parkings) == 0 {
		return nil, &errors.NeedsAParkingSensor{}
	}
	p := new(Place)
	p.id = id
	p.latitude = latitude
	p.longitude = longitude
	p.address = address
	p.freeParkings = 0
	for i := 0; i < len(parkings); i++ {
		if pPlaceID := parkings[i].placeID; pPlaceID != -1 && pPlaceID != p.id {
			return nil, &errors.ParkingSensorIsAlreadyAttached{}
		}
		parkings[i].placeID = p.id

		p.parkings = append(p.parkings, parkings[i].id)

		if status, err := parkings[i].Status(); err == nil && status == free {
			p.freeParkings += 1
		}
	}
	return p, nil
}

// Location returns Place location
func (p *Place) Location() (float64, float64) {
	return p.latitude, p.longitude
}

// Address returns Place address
func (p *Place) Address() string {
	return p.address
}

// Parkings returns Place total parkings and free
func (p *Place) Parkings() (int, int) {
	return len(p.parkings), p.freeParkings
}
