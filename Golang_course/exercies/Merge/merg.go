package merge

// Merge merges two sorted slices into one sorted slice.
func Merge[T any](slice1, slice2 []T, less func(T, T) bool) []T {
	result := make([]T, 0, len(slice1)+len(slice2))
	i, j := 0, 0
	
	for i < len(slice1) && j < len(slice2) {
		if less(slice1[i], slice2[j]) {
			result = append(result, slice1[i])
			i++
		} else {
			result = append(result, slice2[j])
			j++
		}
	}
	
	result = append(result, slice1[i:]...)
	result = append(result, slice2[j:]...)
	
	return result
}
