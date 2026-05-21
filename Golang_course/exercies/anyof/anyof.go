package anyof

// AnyOf returns true if any element in the slice satisfies the predicate.
// Parameters:
// - slice: a slice of any type T
// - predicate: function that returns true/false for a value of type T
// Returns:
// - bool: true if any element satisfies predicate, false otherwise
func AnyOf[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}
