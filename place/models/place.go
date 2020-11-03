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
	return p._latitude, p._longitude
}

// Address returns Place address
func (p *Place) Address() string{
	return p._address
}

// Parkings returns Place total parkings and free
func (p *Place) Parkings() (int, int) {
	return len(p._parkings), p._freeParkings
}
