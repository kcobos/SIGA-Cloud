package model

type ParkingStatus int

const (
	StatusNotValid ParkingStatus = iota - 1
	Free
	Occupied
)

func getStatus(newStatus string) ParkingStatus {
	switch newStatus {
	case "free":
		return Free
	case "occupied":
		return Occupied
	}
	return StatusNotValid
}
