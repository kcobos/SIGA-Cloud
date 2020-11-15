package errors

import (
	"fmt"
)

// StatusNotValid error
type StatusNotValid struct {
}

func (e *StatusNotValid) Error() string {
	return fmt.Sprintf("Status is not valid")
}

// NotInitialized error
type NotInitialized struct {
}

func (e *NotInitialized) Error() string {
	return fmt.Sprintf("Not initialized")
}

// NeedsAParkingSensor error
type NeedsAParkingSensor struct {
}

func (e *NeedsAParkingSensor) Error() string {
	return fmt.Sprintf("A place needs at least one parking sensor")
}

// ParkingSensorIsAlreadyAttached error
type ParkingSensorIsAlreadyAttached struct {
}

func (e *ParkingSensorIsAlreadyAttached) Error() string {
	return fmt.Sprintf("Parking sensor is attached to other place")
}
