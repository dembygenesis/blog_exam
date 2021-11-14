package int_utils

// IntContains checks whether an int exists in a slice of ints
func IntContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
