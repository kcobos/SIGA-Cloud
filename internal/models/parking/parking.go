package parking

import (
	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

// Parking represents a car park
type Parking struct {
	id      int
	status  ParkingStatus
	placeID int
}

// NewParking create a parking object
func NewParking(id int) *Parking {
	p := new(Parking)
	p.id = id
	p.status = statusNotValid
	p.placeID = -1
	return p
}

// ID return the Parking identifier
func (p *Parking) ID() int {
	return p.id
}

// PlaceID return the Parking identifier
func (p *Parking) PlaceID() int {
	return p.placeID
}

// Status returns the current Parking status
func (p *Parking) Status() (ParkingStatus, error) {
	if p.status == statusNotValid {
		return statusNotValid, &errors.StatusNotValid{}
	}
	return p.status, nil
}

// ChangeStatus sets new status to Parking if it's valid
func (p *Parking) ChangeStatus(newStatus string) (bool, error) {
	status := getStatus(newStatus)
	if status > -2 {
		p.status = status
		return true, nil
	}
	return false, &errors.StatusNotValid{}
}
