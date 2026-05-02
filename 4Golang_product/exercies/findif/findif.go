package findif

// FindIf returns the index of the first element that satisfies the predicate.
// Returns -1 if no element satisfies it.
func FindIf[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}
