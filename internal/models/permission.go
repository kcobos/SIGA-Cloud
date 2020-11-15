package model

type userPermission int

const (
	normal   userPermission = 0
	admin                   = 1
	operator                = 2
)
