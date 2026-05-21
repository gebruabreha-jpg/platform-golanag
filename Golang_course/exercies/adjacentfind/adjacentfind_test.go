package adjacentfind

import "testing"

func TestAdjacentFind(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4}
	if idx := AdjacentFind(nums); idx != 2 {
		t.Errorf("Expected 2, got %d", idx)
	}

	nums = []int{1, 2, 3, 4, 5}
	if idx := AdjacentFind(nums); idx != -1 {
		t.Errorf("Expected -1, got %d", idx)
	}
}
