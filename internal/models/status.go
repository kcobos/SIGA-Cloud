package model

type ParkingStatus int

const (
	statusNotValid ParkingStatus = iota - 1
	free
	occupied
)

func getStatus(newStatus string) ParkingStatus {
	switch newStatus {
	case "free":
		return free
	case "occupied":
		return occupied
	}
	return statusNotValid
}
