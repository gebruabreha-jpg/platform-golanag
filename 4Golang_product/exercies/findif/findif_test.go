package findif

import "testing"

func TestFindIf(t *testing.T) {
	nums := []int{1, 3, 5, 8, 9}
	if idx := FindIf(nums, func(n int) bool { return n%2 == 0 }); idx != 3 {
		t.Errorf("Expected 3, got %d", idx)
	}

	nums = []int{1, 3, 5, 7}
	if idx := FindIf(nums, func(n int) bool { return n%2 == 0 }); idx != -1 {
		t.Errorf("Expected -1, got %d", idx)
	}
}
