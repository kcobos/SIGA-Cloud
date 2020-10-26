package parkinghist

import "time"

type parkingHistRecord struct {
	_timestamp time.Time
	_status    int
}

// ParkingHist represents the historical of a car park
type ParkingHist struct {
	_ParkingID  int
	_historical map[int]parkingHistRecord
}
