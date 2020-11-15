package controller

import (
	"github.com/kcobos/SIGA-Cloud/internal/errors"
	. "github.com/kcobos/SIGA-Cloud/internal/models"
)

// Parkings represents a list of Parkings.
type Parkings struct {
	parkingList map[int]*Parking
	lastID      int
}

// NewParkings initialize Parkings struct
// It sets the map and the last ID
func NewParkings() *Parkings {
	p := new(Parkings)
	p.parkingList = make(map[int]*Parking)
	p.lastID = -1
	return p
}

// Len returns length of parking list if it is initialized
func (p *Parkings) Len() (int, error) {
	if p.parkingList == nil {
		return 0, &errors.NotInitialized{}
	}
	return len(p.parkingList), nil
}

// NewParking appends a new parking to the list of parking lots
// Return the new id
func (p *Parkings) NewParking() (int, error) {
	if p.parkingList == nil {
		return -1, &errors.NotInitialized{}
	}

	p.lastID++
	p.parkingList[p.lastID] = NewParking(p.lastID)
	return p.lastID, nil
}
