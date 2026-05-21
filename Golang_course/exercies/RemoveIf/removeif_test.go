package removeif

import "testing"

func TestRemoveIf(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	newLen := RemoveIf(nums, func(n int) bool { return n%2 == 0 })
	
	if newLen != 3 {
		t.Errorf("Expected new length 3, got %d", newLen)
	}
	
	expected := []int{1, 3, 5}
	for i := 0; i < newLen; i++ {
		if nums[i] != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], nums[i])
		}
	}
}
