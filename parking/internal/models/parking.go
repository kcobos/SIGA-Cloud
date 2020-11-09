package parking

import "../errors"

// Parking represents a car park
type Parking struct {
	_ID      int
	_status  parkingStatus
	_placeID int
}

// Status returns current Parking status
func (p *Parking) Status() parkingStatus {
	return p._status
}

// ChangeStatus sets new status to Parking if it's valid
func (p *Parking) ChangeStatus(newStatus string) (bool, error) {
	status := getStatus(newStatus)
	if status > -2 {
		p._status = status
		return true, nil
	}
	return false, &errors.StatusNotValid{}
}
