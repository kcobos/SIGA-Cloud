package model

// User represents a system user
type User struct {
	_username      string
	_permission    userPermission
	_accreditation string
}
