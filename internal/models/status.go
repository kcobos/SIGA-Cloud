package model

type ParkingStatus int

const (
	statusNotValid ParkingStatus = iota - 2
	undefined
	free
	occupied
)

func getStatus(newStatus string) ParkingStatus {
	switch newStatus {
	case "undefined":
		return undefined
	case "free":
		return free
	case "occupied":
		return occupied
	}
	return statusNotValid
}
