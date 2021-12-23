package testutils

import . "supermarket/types"

func Contains(slice []ProduceItem, produce ProduceItem) bool {
	for _, item := range slice {
		if item.Name == produce.Name && item.Code == produce.Code && item.Price == produce.Price {
			return true
		}
	}
	return false
}

// Tests if the passed slices of ProduceItems are equivalent.
// Returns true if they are and an error message if they aren't.
func Equivalent(first []ProduceItem, second []ProduceItem) (bool, string) {
	if len(first) != len(second) {
		return false, "lengths differ"
	}
	for _, item := range first {
		if !Contains(second, item) {
			return false, "contents differ"
		}
	}
	for _, item := range second {
		if !Contains(first, item) {
			return false, "contents differ"
		}
	}
	return true, ""
}
