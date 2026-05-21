package issorted

// IsSorted checks if the slice is sorted in ascending order.
func IsSorted[T comparable](slice []T, less func(T, T) bool) bool {
	for i := 0; i < len(slice)-1; i++ {
		if less(slice[i+1], slice[i]) {
			return false
		}
	}
	return true
}
