package tests

import (
	"testing"
)

// TestMapDeclaration tests basic map declarations
func TestMapDeclaration(t *testing.T) {
	var mapa1 map[int]complex128
	
	if mapa1 != nil {
		t.Error("Uninitialized map should be nil")
	}
}

// TestMapNil tests that empty maps and slices are nil
func TestMapNil(t *testing.T) {
	var s []int
	var mapa1 map[int]complex128
	
	if s != nil {
		t.Error("Uninitialized slice should be nil")
	}
	if mapa1 != nil {
		t.Error("Uninitialized map should be nil")
	}
}

// TestMapCreation tests creating maps with make
func TestMapCreation(t *testing.T) {
	mapa := make(map[int]complex128)
	
	if mapa == nil {
		t.Error("Map created with make should not be nil")
	}
	if len(mapa) != 0 {
		t.Errorf("Newly created map should have length 0, got %d", len(mapa))
	}
	
	// Test map literal syntax
	mapa2 := map[int]complex128{}
	if mapa2 == nil {
		t.Error("Map created with literal syntax should not be nil")
	}
}

// TestMapLiteralInitialization tests map initialization with literals
func TestMapLiteralInitialization(t *testing.T) {
	demografía := map[string]int{
		"España":   46940000,
		"Portugal": 10280000,
		"Francia":  66990000,
	}
	
	if len(demografía) != 3 {
		t.Errorf("Expected map length 3, got %d", len(demografía))
	}
	
	if demografía["España"] != 46940000 {
		t.Errorf("Expected 46940000, got %d", demografía["España"])
	}
	if demografía["Portugal"] != 10280000 {
		t.Errorf("Expected 10280000, got %d", demografía["Portugal"])
	}
	if demografía["Francia"] != 66990000 {
		t.Errorf("Expected 66990000, got %d", demografía["Francia"])
	}
}

// TestMapReadWrite tests map read and write operations
func TestMapReadWrite(t *testing.T) {
	var demografía map[string]int = make(map[string]int)
	
	demografía["España"] = 46940000
	demografía["Portugal"] = 10280000
	demografía["Francia"] = 66990000
	
	if demografía["España"] != 46940000 {
		t.Error("Failed to write/read map entry")
	}
	
	// Test updating an entry
	demografía["España"] = 47000000
	if demografía["España"] != 47000000 {
		t.Error("Failed to update map entry")
	}
}

// TestMapDefaultValue tests default value for missing keys
func TestMapDefaultValue(t *testing.T) {
	demografía := map[string]int{
		"España":   46940000,
		"Portugal": 10280000,
	}
	
	// Accessing non-existent key returns zero value
	italia := demografía["Italia"]
	if italia != 0 {
		t.Errorf("Expected default value 0 for missing key, got %d", italia)
	}
}

// TestMapValueCheck tests checking if a key exists
func TestMapValueCheck(t *testing.T) {
	demografía := map[string]int{
		"España":   46940000,
		"Portugal": 10280000,
	}
	
	// Check for existing key
	if value, ok := demografía["España"]; !ok {
		t.Error("Key 'España' should exist")
	} else if value != 46940000 {
		t.Errorf("Expected 46940000, got %d", value)
	}
	
	// Check for non-existing key
	if value, ok := demografía["Italia"]; ok {
		t.Error("Key 'Italia' should not exist")
	} else if value != 0 {
		t.Errorf("Expected default value 0, got %d", value)
	}
}

// TestMapDelete tests deleting entries from a map
func TestMapDelete(t *testing.T) {
	demografía := map[string]int{
		"España":   46940000,
		"Portugal": 10280000,
		"Francia":  66990000,
	}
	
	delete(demografía, "España")
	
	if _, ok := demografía["España"]; ok {
		t.Error("Key 'España' should have been deleted")
	}
	if len(demografía) != 2 {
		t.Errorf("Expected map length 2 after deletion, got %d", len(demografía))
	}
	
	// Deleting non-existent key should not cause error
	delete(demografía, "Italia")
	if len(demografía) != 2 {
		t.Error("Deleting non-existent key should not change map length")
	}
}

// TestMapClear tests the clear function for maps (Go 1.21+)
func TestMapClear(t *testing.T) {
	a := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: false,
		5: true,
		6: false,
	}
	
	if len(a) != 6 {
		t.Errorf("Expected initial length 6, got %d", len(a))
	}
	
	clear(a)
	
	if len(a) != 0 {
		t.Errorf("Expected length 0 after clear, got %d", len(a))
	}
	
	// Accessing cleared map returns default value
	if a[3] != false {
		t.Error("Accessing cleared map should return default value")
	}
}

// TestMapIteration tests iterating over maps
func TestMapIteration(t *testing.T) {
	demografía := map[string]int{
		"España":   46940000,
		"Portugal": 10280000,
		"Francia":  66990000,
	}
	
	count := 0
	total := 0
	for k, v := range demografía {
		count++
		total += v
		if k == "" {
			t.Error("Key should not be empty")
		}
		if v <= 0 {
			t.Error("Value should be positive")
		}
	}
	
	if count != 3 {
		t.Errorf("Expected to iterate 3 times, iterated %d times", count)
	}
	
	expectedTotal := 46940000 + 10280000 + 66990000
	if total != expectedTotal {
		t.Errorf("Expected total %d, got %d", expectedTotal, total)
	}
}

// TestMapKeyTypes tests that map keys must be comparable
func TestMapKeyTypes(t *testing.T) {
	// Test with various key types
	intMap := map[int]string{1: "one", 2: "two"}
	if intMap[1] != "one" {
		t.Error("Int keys should work")
	}
	
	stringMap := map[string]int{"one": 1, "two": 2}
	if stringMap["one"] != 1 {
		t.Error("String keys should work")
	}
	
	type Point struct {
		X, Y int
	}
	structMap := map[Point]string{{0, 0}: "origin", {1, 1}: "diagonal"}
	if structMap[Point{0, 0}] != "origin" {
		t.Error("Struct keys should work")
	}
}

// TestMapValueTypes tests maps with various value types
func TestMapValueTypes(t *testing.T) {
	// Map with slice values
	sliceMap := map[string][]int{
		"primes": {2, 3, 5, 7},
		"evens":  {2, 4, 6, 8},
	}
	if len(sliceMap["primes"]) != 4 {
		t.Error("Map with slice values should work")
	}
	
	// Map with map values
	nestedMap := map[string]map[string]int{
		"scores": {"alice": 100, "bob": 85},
	}
	if nestedMap["scores"]["alice"] != 100 {
		t.Error("Map with map values should work")
	}
}

// TestMapCapacityHint tests creating maps with capacity hint
func TestMapCapacityHint(t *testing.T) {
	// Create map with capacity hint
	mapa := make(map[int]string, 100)
	
	if mapa == nil {
		t.Error("Map with capacity hint should not be nil")
	}
	if len(mapa) != 0 {
		t.Error("Initial length should be 0")
	}
	
	// Add elements beyond initial capacity (should work fine)
	for i := 0; i < 150; i++ {
		mapa[i] = "value"
	}
	
	if len(mapa) != 150 {
		t.Errorf("Expected length 150, got %d", len(mapa))
	}
}

// TestMapZeroValue tests zero value behavior
func TestMapZeroValue(t *testing.T) {
	var nilMap map[string]int
	
	// Reading from nil map returns zero value
	value := nilMap["key"]
	if value != 0 {
		t.Errorf("Expected 0, got %d", value)
	}
	
	// Checking existence in nil map
	_, ok := nilMap["key"]
	if ok {
		t.Error("Key should not exist in nil map")
	}
	
	// Length of nil map is 0
	if len(nilMap) != 0 {
		t.Errorf("Expected length 0, got %d", len(nilMap))
	}
}

// TestMapModification tests modifying maps during iteration
func TestMapModification(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	
	// Add elements during iteration (not recommended but should work)
	seen := make(map[string]bool)
	for k := range m {
		seen[k] = true
	}
	
	if len(seen) != 3 {
		t.Errorf("Expected to see 3 keys, saw %d", len(seen))
	}
}
