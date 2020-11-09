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
	return fmt.Sprintf("Parkings not initialized")
}