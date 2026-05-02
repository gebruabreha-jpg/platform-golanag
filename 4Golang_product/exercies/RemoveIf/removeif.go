package removeif

// RemoveIf removes all elements that satisfy the predicate.
// Returns the new length of the slice after removal.
func RemoveIf[T any](slice []T, predicate func(T) bool) int {
	j := 0
	for i := range slice {
		if !predicate(slice[i]) {
			slice[j] = slice[i]
			j++
		}
	}
	return j
}
