package tests

import (
	"reflect"
	"testing"
	"unicode"
)

// TestArrayDeclaration tests basic array declarations and initialization
func TestArrayDeclaration(t *testing.T) {
	// Test array with const size (valid)
	const tamaño = 10
	var array [tamaño]int
	if len(array) != 10 {
		t.Errorf("Expected array length 10, got %d", len(array))
	}

	// Test array initialization
	array2 := [2]float64{}
	array3 := [3]float64{}
	if reflect.TypeOf(array2) == reflect.TypeOf(array3) {
		t.Error("Arrays with different sizes should have different types")
	}
}

// TestArrayComparison tests array comparison
func TestArrayComparison(t *testing.T) {
	array1 := [2]float64{}
	array2 := [2]float64{}
	if array1 != array2 {
		t.Error("Empty arrays of same type should be equal")
	}
	
	array1[0] = 1.0
	if array1 == array2 {
		t.Error("Arrays with different values should not be equal")
	}
}

// TestArrayTypeSignature tests that array size is part of type signature
func TestArrayTypeSignature(t *testing.T) {
	array1 := [10]int{}
	array2 := [5]int{}
	array3 := [10]complex128{}
	
	type1 := reflect.TypeOf(array1).String()
	type2 := reflect.TypeOf(array2).String()
	type3 := reflect.TypeOf(array3).String()
	
	if type1 != "[10]int" {
		t.Errorf("Expected type [10]int, got %s", type1)
	}
	if type2 != "[5]int" {
		t.Errorf("Expected type [5]int, got %s", type2)
	}
	if type3 != "[10]complex128" {
		t.Errorf("Expected type [10]complex128, got %s", type3)
	}
}

// TestArrayOperators tests array access and slicing operators
func TestArrayOperators(t *testing.T) {
	// Test random access
	array := [3]int{1, 2, 3}
	if array[0] != 1 || array[1] != 2 || array[2] != 3 {
		t.Error("Array random access failed")
	}
	
	// Test write access
	array[1] = 10
	if array[1] != 10 {
		t.Error("Array write access failed")
	}
	
	// Test slicing
	slice := array[0:2]
	if len(slice) != 2 || slice[0] != 1 || slice[1] != 10 {
		t.Error("Array slicing failed")
	}
}

// TestArrayInitialization tests various array initialization methods
func TestArrayInitialization(t *testing.T) {
	// Literal initialization
	a := [2]int{2, 4}
	if a[0] != 2 || a[1] != 4 {
		t.Error("Literal initialization failed")
	}
	
	// Partial initialization
	c := [3]byte{'g', 'o'}
	if c[0] != 'g' || c[1] != 'o' || c[2] != 0 {
		t.Error("Partial initialization failed, expected default value 0 for uninitialized element")
	}
	
	// Auto-size with ...
	d := [...]byte{'g', 'o'}
	if len(d) != 2 {
		t.Errorf("Auto-size array should have length 2, got %d", len(d))
	}
	
	// Selective initialization
	cuadrados := [...]int{1: 1, 2: 4, 3: 9, 4: 16}
	if cuadrados[0] != 0 || cuadrados[1] != 1 || cuadrados[4] != 16 {
		t.Error("Selective initialization failed")
	}
}

// TestArrayIndexedInitialization tests indexed initialization with iota
func TestArrayIndexedInitialization(t *testing.T) {
	const (
		España = iota
		Portugal
		Francia
	)
	capitales := [...]string{
		España:   "Madrid",
		Portugal: "Lisboa",
		Francia:  "París",
	}
	
	if capitales[España] != "Madrid" {
		t.Errorf("Expected Madrid, got %s", capitales[España])
	}
	if capitales[Portugal] != "Lisboa" {
		t.Errorf("Expected Lisboa, got %s", capitales[Portugal])
	}
	if capitales[Francia] != "París" {
		t.Errorf("Expected París, got %s", capitales[Francia])
	}
}

// TestArrayMutability tests that arrays are mutable
func TestArrayMutability(t *testing.T) {
	array := [3]int{1, 2, 3}
	array[1] = 10
	if array[1] != 10 {
		t.Error("Arrays should be mutable")
	}
}

// capitalizeByValue does not modify the array (pass by value)
func capitalizeByValue(cadena [2]byte) {
	cadena[0] = byte(unicode.ToUpper(rune(cadena[0])))
}

// capitalizeByReference modifies the array (pass by reference)
func capitalizeByReference(cadena *[2]byte) {
	(*cadena)[0] = byte(unicode.ToUpper(rune(cadena[0])))
}

// TestArrayPassByValue tests that arrays are passed by value
func TestArrayPassByValue(t *testing.T) {
	array := [...]byte{'g', 'o'}
	
	capitalizeByValue(array)
	if array[0] != 'g' {
		t.Error("Array should not be modified when passed by value")
	}
	
	capitalizeByReference(&array)
	if array[0] != 'G' {
		t.Error("Array should be modified when passed by reference")
	}
}

// TestArrayToSliceConversion tests converting array to slice with [:]
func TestArrayToSliceConversion(t *testing.T) {
	array := [...]byte{'g', 'o'}
	
	// Test that array[:] creates a slice
	slice := array[:]
	sliceType := reflect.TypeOf(slice).String()
	arrayType := reflect.TypeOf(array).String()
	
	if sliceType != "[]uint8" {
		t.Errorf("Expected slice type []uint8, got %s", sliceType)
	}
	if arrayType != "[2]uint8" {
		t.Errorf("Expected array type [2]uint8, got %s", arrayType)
	}
	
	// Test string conversion
	str := string(array[:])
	if str != "go" {
		t.Errorf("Expected 'go', got '%s'", str)
	}
}

// TestArrayDefaultValues tests default values in arrays
func TestArrayDefaultValues(t *testing.T) {
	var intArray [3]int
	for i := 0; i < 3; i++ {
		if intArray[i] != 0 {
			t.Errorf("Default value for int should be 0, got %d", intArray[i])
		}
	}
	
	var floatArray [2]float64
	for i := 0; i < 2; i++ {
		if floatArray[i] != 0.0 {
			t.Errorf("Default value for float64 should be 0.0, got %f", floatArray[i])
		}
	}
	
	var stringArray [2]string
	for i := 0; i < 2; i++ {
		if stringArray[i] != "" {
			t.Errorf("Default value for string should be empty, got '%s'", stringArray[i])
		}
	}
}

// TestArrayFibonacci tests array with Fibonacci sequence
func TestArrayFibonacci(t *testing.T) {
	fibonacci := [11]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	
	// Verify first two elements
	if fibonacci[0] != 1 || fibonacci[1] != 1 {
		t.Error("First two Fibonacci numbers should be 1, 1")
	}
	
	// Verify Fibonacci property for remaining elements
	for i := 2; i < 11; i++ {
		if fibonacci[i] != fibonacci[i-1]+fibonacci[i-2] {
			t.Errorf("Fibonacci property violated at index %d", i)
		}
	}
}
