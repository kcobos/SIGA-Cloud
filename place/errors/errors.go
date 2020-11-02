package errors

import (
	"fmt"
)

// NotInitialized error
type NotInitialized struct {
}
func (e *NotInitialized) Error() string {
	return fmt.Sprintf("Places not initialized")
}