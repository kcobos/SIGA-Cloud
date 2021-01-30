package model

import (
	"container/list"
	"time"

	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

// Parking represents a car park
type Parking struct {
	id      int
	status  ParkingStatus
	placeID int
	history *list.List
}

// NewParking create a parking object
func NewParking(id int) *Parking {
	p := new(Parking)
	p.id = id
	p.status = statusNotValid
	p.placeID = -1
	p.history = list.New()
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
	if status >= free && status <= occupied {
		p.status = status
		psh := new(ParkingStatusHistory)
		psh.status = status
		psh.timestamp = time.Now().UTC().Unix()
		p.history.PushFront(psh)
		return true, nil
	}
	return false, &errors.StatusNotValid{}
}
