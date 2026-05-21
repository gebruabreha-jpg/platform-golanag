package tests

import (
	"reflect"
	"testing"
)

// TestSliceDeclaration tests basic slice declarations
func TestSliceDeclaration(t *testing.T) {
	// Test empty slice initialization
	sa := []int{}
	if sa == nil {
		t.Error("Empty slice literal should not be nil")
	}
	if len(sa) != 0 {
		t.Errorf("Empty slice should have length 0, got %d", len(sa))
	}
	
	// Test nil slice
	var slice []int
	if slice != nil {
		t.Error("Uninitialized slice should be nil")
	}
}

// TestSliceStructure tests slice pointer, len, and cap
func TestSliceStructure(t *testing.T) {
	s := make([]int, 5, 10)
	
	if len(s) != 5 {
		t.Errorf("Expected length 5, got %d", len(s))
	}
	if cap(s) != 10 {
		t.Errorf("Expected capacity 10, got %d", cap(s))
	}
}

// lessThan compares two byte slices
func lessThan(slice1, slice2 []byte) int {
	items := len(slice1)
	if len(slice2) < len(slice1) {
		items = len(slice2)
	}
	for idx := 0; idx < items; idx++ {
		if slice1[idx] < slice2[idx] {
			return -1
		}
		if slice1[idx] > slice2[idx] {
			return +1
		}
	}
	if items == len(slice1) && items == len(slice2) {
		return 0
	}
	if items == len(slice1) {
		return -1
	}
	return +1
}

// TestSliceComparison tests slice comparison function
func TestSliceComparison(t *testing.T) {
	slice1 := []byte{'g', 'o'}
	slice2 := []byte{'g', 'o'}
	slice3 := []byte{'g', 'o', '!'}
	slice4 := []byte{'h', 'i'}
	
	if lessThan(slice1, slice2) != 0 {
		t.Error("Equal slices should return 0")
	}
	if lessThan(slice1, slice3) != -1 {
		t.Error("Shorter slice should be less than longer slice with same prefix")
	}
	if lessThan(slice3, slice1) != +1 {
		t.Error("Longer slice should be greater than shorter slice with same prefix")
	}
	if lessThan(slice1, slice4) != -1 {
		t.Error("'go' should be less than 'hi'")
	}
}

// TestSliceAccess tests slice random access operator
func TestSliceAccess(t *testing.T) {
	slice := []int{1, 2, 3}
	
	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 {
		t.Error("Slice random access failed")
	}
	
	slice[1] = 10
	if slice[1] != 10 {
		t.Error("Slice write access failed")
	}
}

// TestSliceArrayConversion tests converting array to slice
func TestSliceArrayConversion(t *testing.T) {
	slice1 := []byte{'g', 'o'}
	var array2 = [2]byte{103, 111}
	
	result := lessThan(slice1, array2[:])
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

// TestRuneSliceConversion tests converting string to rune slice
func TestRuneSliceConversion(t *testing.T) {
	cadena := "sin(θ)=cos(π/2-θ)"
	runas := []rune(cadena)
	
	if len(runas) == 0 {
		t.Error("Rune slice should not be empty")
	}
	
	// Check first few characters
	expected := []rune{'s', 'i', 'n', '(', 'θ'}
	for i, r := range expected {
		if i < len(runas) && runas[i] != r {
			t.Errorf("Expected rune %v at position %d, got %v", r, i, runas[i])
		}
	}
}

// TestSliceWithMake tests creating slices with make
func TestSliceWithMake(t *testing.T) {
	s := make([]float64, 1)
	
	if len(s) != 1 {
		t.Errorf("Expected length 1, got %d", len(s))
	}
	
	s[0] = 0.0
	if s[0] != 0.0 {
		t.Error("Failed to set slice element")
	}
	
	// Test that accessing beyond length panics
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when accessing beyond slice length")
		}
	}()
	_ = s[1] // This should panic
}

// TestSliceCapacityExceeding tests behavior when slice capacity is exceeded
func TestSliceCapacityExceeding(t *testing.T) {
	s1 := make([]bool, 0, 2)
	
	if len(s1) != 0 {
		t.Errorf("Expected length 0, got %d", len(s1))
	}
	if cap(s1) != 2 {
		t.Errorf("Expected capacity 2, got %d", cap(s1))
	}
	
	// Test that accessing beyond length panics even with available capacity
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when accessing beyond slice length")
		}
	}()
	_ = s1[0] // This should panic
}

// TestSliceDefaultValues tests default values in slices
func TestSliceDefaultValues(t *testing.T) {
	s1 := make([]bool, 2)
	if s1[0] != false || s1[1] != false {
		t.Error("Default value for bool should be false")
	}
	
	s2 := make([]int, 3)
	for i := 0; i < 3; i++ {
		if s2[i] != 0 {
			t.Errorf("Default value for int should be 0, got %d", s2[i])
		}
	}
	
	s3 := make([]string, 4)
	for i := 0; i < 4; i++ {
		if s3[i] != "" {
			t.Errorf("Default value for string should be empty, got '%s'", s3[i])
		}
	}
}

// TestSliceAppend tests the append function
func TestSliceAppend(t *testing.T) {
	slice := []int{}
	slice = append(slice, 1)
	
	if len(slice) != 1 || slice[0] != 1 {
		t.Error("Append failed")
	}
	
	slice = append(slice, 2, 3, 4)
	if len(slice) != 4 {
		t.Errorf("Expected length 4, got %d", len(slice))
	}
	
	// Test appending a slice
	slice2 := []int{5, 6}
	slice = append(slice, slice2...)
	if len(slice) != 6 {
		t.Errorf("Expected length 6, got %d", len(slice))
	}
}

// updateValue modifies slice in-place
func updateValue(s []int, value int) {
	if len(s) > 0 {
		s[0] = value
	}
}

// TestSlicePassByValue tests slice modification when passed by value
func TestSlicePassByValue(t *testing.T) {
	slice := []int{1, 2, 3}
	updateValue(slice, 4)
	
	if slice[0] != 4 {
		t.Error("Slice should be modified in-place even when passed by value")
	}
}

// addValue tries to append to a slice (doesn't work with pass by value)
func addValue(s []int, value int) {
	s = append(s, value)
}

// TestSliceAppendByValue tests that append doesn't work with pass by value
func TestSliceAppendByValue(t *testing.T) {
	slice := []int{}
	addValue(slice, 1)
	
	if len(slice) != 0 {
		t.Error("Slice should not be modified when append is used in pass by value")
	}
}

// addValueByReference appends to a slice by reference
func addValueByReference(s *[]int, value int) {
	*s = append(*s, value)
}

// TestSliceAppendByReference tests append with pass by reference
func TestSliceAppendByReference(t *testing.T) {
	slice := []int{}
	addValueByReference(&slice, 1)
	
	if len(slice) != 1 || slice[0] != 1 {
		t.Error("Slice should be modified when passed by reference")
	}
	
	slice = make([]int, 0, 1)
	addValueByReference(&slice, 2)
	
	if len(slice) != 1 || slice[0] != 2 {
		t.Error("Slice should be modified when passed by reference")
	}
}

// TestSliceCopy tests the copy function
func TestSliceCopy(t *testing.T) {
	s2 := make([]int, 3)
	s2[1] = 1
	
	s3 := make([]int, 3)
	copied := copy(s3, s2)
	
	if copied != 3 {
		t.Errorf("Expected to copy 3 elements, copied %d", copied)
	}
	
	s3[1] = 2
	if s2[1] != 1 {
		t.Error("Original slice should not be affected by copy")
	}
	if s3[1] != 2 {
		t.Error("Copied slice should be independent")
	}
}

// TestSliceCopyMinimum tests that copy copies minimum of src and dst lengths
func TestSliceCopyMinimum(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{}
	
	copied := copy(s1, s2)
	if copied != 0 {
		t.Errorf("Expected to copy 0 elements, copied %d", copied)
	}
	
	copied = copy(s2, s1)
	if copied != 0 {
		t.Errorf("Expected to copy 0 elements, copied %d", copied)
	}
	
	s3 := make([]int, 0, 10)
	copied = copy(s3, s1)
	if copied != 0 {
		t.Errorf("Expected to copy 0 elements (capacity doesn't matter), copied %d", copied)
	}
}

// TestSliceClear tests the clear function (Go 1.21+)
func TestSliceClear(t *testing.T) {
	z := []int{0, 1, 2, 3, 4, 5, 6, 7}
	
	originalLen := len(z)
	originalCap := cap(z)
	
	clear(z)
	
	if len(z) != originalLen {
		t.Errorf("Clear should not change length, expected %d, got %d", originalLen, len(z))
	}
	if cap(z) != originalCap {
		t.Errorf("Clear should not change capacity, expected %d, got %d", originalCap, cap(z))
	}
	
	for i, v := range z {
		if v != 0 {
			t.Errorf("Element at index %d should be 0 after clear, got %d", i, v)
		}
	}
}

// TestSliceLiteralInitialization tests slice literal initialization
func TestSliceLiteralInitialization(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	
	if len(slice) != 5 {
		t.Errorf("Expected length 5, got %d", len(slice))
	}
	
	for i, expected := range []int{1, 2, 3, 4, 5} {
		if slice[i] != expected {
			t.Errorf("Expected %d at index %d, got %d", expected, i, slice[i])
		}
	}
}

// TestSliceSubslicing tests creating sub-slices
func TestSliceSubslicing(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	
	sub := slice[2:4]
	if len(sub) != 2 || sub[0] != 2 || sub[1] != 3 {
		t.Error("Sub-slice failed")
	}
	
	// Test that modifying sub-slice affects original
	sub[0] = 99
	if slice[2] != 99 {
		t.Error("Sub-slice should share underlying array with original")
	}
}

// TestSliceNilCheck tests checking for nil slices
func TestSliceNilCheck(t *testing.T) {
	var nilSlice []int
	emptySlice := []int{}
	
	if nilSlice != nil {
		t.Error("Uninitialized slice should be nil")
	}
	if emptySlice == nil {
		t.Error("Empty slice literal should not be nil")
	}
	
	// Both have length 0
	if len(nilSlice) != 0 || len(emptySlice) != 0 {
		t.Error("Both nil and empty slices should have length 0")
	}
}

// TestSliceTypeConversion tests slice type is distinct from array type
func TestSliceTypeConversion(t *testing.T) {
	slice := []int{1, 2, 3}
	sliceType := reflect.TypeOf(slice).String()
	
	if sliceType != "[]int" {
		t.Errorf("Expected slice type []int, got %s", sliceType)
	}
	
	array := [3]int{1, 2, 3}
	arrayType := reflect.TypeOf(array).String()
	
	if arrayType != "[3]int" {
		t.Errorf("Expected array type [3]int, got %s", arrayType)
	}
}
