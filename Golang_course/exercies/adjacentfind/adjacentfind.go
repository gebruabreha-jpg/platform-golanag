package adjacentfind

// AdjacentFind returns the index of the first pair of adjacent equal elements.
// Returns -1 if no such pair exists.
func AdjacentFind[T comparable](slice []T) int {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == slice[i+1] {
			return i
		}
	}
	return -1
}
