package place

// Place represents a parking location
type Place struct {
	_ID                   int
	_latitude, _longitude float64
	_address              string
	_parkings             []int
	_freeParkings         int
}

// Coordinates returns Place location
func (p *Place) Coordinates() (float64, float64) {

}

// AddParking adds a Parking to a Place
func (p *Place) AddParking(parkingID int) bool {

}
