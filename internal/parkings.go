package controller

import (
	"github.com/kcobos/SIGA-Cloud/internal/errors"
	. "github.com/kcobos/SIGA-Cloud/internal/models/parking"
)

// Parkings represents a list of Parkings.
type Parkings struct {
	parkingList map[int]Parking
	lastID      int
}

// NewParkings initialize Parkings struct
// It sets the map and the last ID
func NewParkings() *Parkings {
	p := new(Parkings)
	p.parkingList = make(map[int]Parking)
	p.lastID = 0
	return p
}

// Len returns length of parking list if it is initialized
func (p *Parkings) Len() (int, error) {
	if p.parkingList == nil {
		return 0, &errors.NotInitialized{}
	}
	return len(p.parkingList), nil
}
