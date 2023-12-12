package oxylabs

// Checks if the parameter is in the list of accepted parameters.
func InList[T comparable](val T, list []T) bool {
	for _, item := range list {
		if item == val {
			return true
		}
	}
	return false
}