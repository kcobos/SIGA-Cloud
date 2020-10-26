package parking

// Parking represents a car park
type Parking struct {
	_ID      int
	_status  parkingStatus
	_placeID int
}

// Status returns current Parking status
func (p *Parking) Status() string {

}

// ChangeStatus sets new status to Parking if it's valid
func (p *Parking) ChangeStatus(newStatus string) bool {

}

// AddToPlace attaches a Place to a Parking
func (p *Parking) AddToPlace(placeID int) {

}
