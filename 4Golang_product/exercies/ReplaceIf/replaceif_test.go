package replaceif

import "testing"

func TestReplaceIf(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	count := ReplaceIf(nums, 0, func(n int) bool { return n%2 == 0 })
	
	if count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
	
	expected := []int{1, 0, 3, 0, 5}
	for i := range nums {
		if nums[i] != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], nums[i])
		}
	}
}
