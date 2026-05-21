package replaceif

// ReplaceIf replaces all elements that satisfy the predicate with newValue.
// Returns the count of replaced elements.
func ReplaceIf[T any](slice []T, newValue T, predicate func(T) bool) int {
	count := 0
	for i, v := range slice {
		if predicate(v) {
			slice[i] = newValue
			count++
		}
	}
	return count
}
