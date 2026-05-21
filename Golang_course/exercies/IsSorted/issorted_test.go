package issorted

import "testing"

func TestIsSorted(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	if !IsSorted(nums, func(a, b int) bool { return a < b }) {
		t.Error("Expected true, got false")
	}

	nums = []int{1, 3, 2, 4, 5}
	if IsSorted(nums, func(a, b int) bool { return a < b }) {
		t.Error("Expected false, got true")
	}

	nums = []int{5, 4, 3, 2, 1}
	if IsSorted(nums, func(a, b int) bool { return a < b }) {
		t.Error("Expected false, got true")
	}
}
