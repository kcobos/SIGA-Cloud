package model

// ParkingStatusHistory represents the Parking's status in a specific time
type ParkingStatusHistory struct {
	status    ParkingStatus
	timestamp int64
}
