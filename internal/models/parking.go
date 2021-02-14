package model

import (
	"time"

	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

// Parking represents a car park
type Parking struct {
	ID      int64
	Status  ParkingStatus
	PlaceID int64
	History []ParkingStatusHistory `pg:"rel:has-many"`
}

// NewParking create a parking object
func NewParking() *Parking {
	p := new(Parking)
	p.Status = StatusNotValid
	p.PlaceID = -1
	return p
}

// GetStatus returns the current Parking status
func (p *Parking) GetStatus() (ParkingStatus, error) {
	if p.Status == StatusNotValid {
		return StatusNotValid, &errors.StatusNotValid{}
	}
	return p.Status, nil
}

// ChangeStatus sets new status to Parking if it's valid
func (p *Parking) ChangeStatus(newStatus string) error {
	status := getStatus(newStatus)
	if status >= Free && status <= Occupied {
		p.Status = status
		psh := ParkingStatusHistory{
			ParkingID: p.ID,
			Status:    status,
			Timestamp: time.Now().UTC().Unix(),
		}
		p.History = append(p.History, psh)
		return nil
	}
	return &errors.StatusNotValid{}
}
