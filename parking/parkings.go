package parkings

import (
	. "./models"
	"./errors"
)

// Parkings represents a list of Parkings.
type Parkings struct {
	parkingList map[int]Parking
	lastID      int
}

// Init initialize Parkings struct.
// It sets the map and the last ID
func (p *Parkings) Init() {
	p.parkingList = make(map[int]Parking)
	p.lastID = 0
}

// Len returns length of parking list if it is initialized
func (p *Parkings) Len() (int, error) {
	if p.parkingList == nil {
		return 0, &errors.NotInitialized{}
	}
	return len(p.parkingList), nil
}
