package model

// ParkingStatusHistory represents the Parking's status in a specific time
type ParkingStatusHistory struct {
	ID        int
	ParkingID int64
	Status    ParkingStatus
	Timestamp int64
}
