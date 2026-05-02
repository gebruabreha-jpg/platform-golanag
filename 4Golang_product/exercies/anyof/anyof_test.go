package anyof

import "testing"

func TestAnyOf(t *testing.T) {
	nums := []int{1, 3, 5, 7}
	if AnyOf(nums, func(n int) bool { return n%2 == 0 }) {
		t.Error("Expected false, got true")
	}

	nums = []int{1, 2, 3, 4}
	if !AnyOf(nums, func(n int) bool { return n%2 == 0 }) {
		t.Error("Expected true, got false")
	}
}
