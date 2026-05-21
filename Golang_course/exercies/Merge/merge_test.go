package merge

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1, 3, 5}
	nums2 := []int{2, 4, 6}
	result := Merge(nums1, nums2, func(a, b int) bool { return a < b })
	
	expected := []int{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], result[i])
		}
	}
}
