package parking

type parkingStatus int

const (
	statusNotValid parkingStatus = iota -2
	undefined
	free
	occupied
)

func getStatus(newStatus string) parkingStatus {
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