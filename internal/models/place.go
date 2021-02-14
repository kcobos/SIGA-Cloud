package model

import (
	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

// Place represents a parking location
type Place struct {
	ID                  int64
	Latitude, Longitude float64
	Address             string
	Parkings            []*Parking `pg:"rel:has-many"`
	FreeParkings        int
}

// NewPlace create a place object
func NewPlace(latitude, longitude float64, address string, parkings []*Parking) (*Place, error) {
	if len(parkings) == 0 {
		return nil, &errors.NeedsAParkingSensor{}
	}
	p := new(Place)
	p.Latitude = latitude
	p.Longitude = longitude
	p.Address = address
	p.FreeParkings = 0
	for i := 0; i < len(parkings); i++ {
		if pPlaceID := parkings[i].PlaceID; pPlaceID != -1 {
			return nil, &errors.ParkingSensorIsAlreadyAttached{}
		}
		parkings[i].PlaceID = p.ID

		p.Parkings = append(p.Parkings, parkings[i])

		if status, err := parkings[i].GetStatus(); err == nil && status == Free {
			p.FreeParkings++
		}
	}
	return p, nil
}

// Location returns Place location
func (p *Place) Location() (float64, float64) {
	return p.Latitude, p.Longitude
}

// GetParkings returns Place total parkings and free
func (p *Place) GetParkings() (int, int) {
	return len(p.Parkings), p.FreeParkings
}
