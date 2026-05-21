package equal

import "testing"

func TestEqual(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{1, 2, 3, 4}
	if !Equal(nums1, nums2) {
		t.Error("Expected true, got false")
	}

	nums3 := []int{1, 2, 3}
	if Equal(nums1, nums3) {
		t.Error("Expected false, got true")
	}

	nums4 := []int{1, 2, 3, 5}
	if Equal(nums1, nums4) {
		t.Error("Expected false, got true")
	}
}
