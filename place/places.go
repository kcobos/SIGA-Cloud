package places

import (
	. "./models"
	"./errors"
)

// Places represents a list of Places.
type Places struct {
	placeList map[int]Place
	lastID      int
}

// Init initialize Places struct.
// It sets the map and the last ID
func (p *Places) Init() {
	p.placeList = make(map[int]Place)
	p.lastID = 0
}

// Len returns length of place list if it is initialized
func (p *Places) Len() (int, error) {
	if p.placeList == nil {
		return 0, &errors.NotInitialized{}
	}
	return len(p.placeList), nil
}
