package tests

import (
	"errors"
	"fmt"
	"testing"
)

const B = 1000

// Bucket operations for sorting integers

// InitBucket initializes a bucket with B+1 entries
func InitBucket() [][]int {
	bucket := make([][]int, 1+B)
	for i := 0; i <= B; i++ {
		bucket[i] = []int{}
	}
	return bucket
}

// Insert adds number i to the i-th bucket
func Insert(i int, bucket [][]int) error {
	if i >= len(bucket) {
		return fmt.Errorf("item out of bounds: %v > %v", i, B)
	}
	bucket[i] = append(bucket[i], i)
	return nil
}

// Pop returns all items in ascending order
func Pop(bucket [][]int) []int {
	result := []int{}
	for i := 0; i <= B; i++ {
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)
		}
	}
	return result
}

// BucketSort sorts all numbers using buckets
func BucketSort(items []int) error {
	bckt := InitBucket()
	
	for _, v := range items {
		if err := Insert(v, bckt); err != nil {
			return err
		}
	}
	
	copy(items, Pop(bckt))
	return nil
}

// TestInitBucket tests bucket initialization
func TestInitBucket(t *testing.T) {
	bucket := InitBucket()
	
	if len(bucket) != B+1 {
		t.Errorf("Expected bucket length %d, got %d", B+1, len(bucket))
	}
	
	for i := 0; i <= B; i++ {
		if bucket[i] == nil {
			t.Errorf("Bucket at index %d should not be nil", i)
		}
		if len(bucket[i]) != 0 {
			t.Errorf("Bucket at index %d should be empty", i)
		}
	}
}

// TestInsert tests inserting items into buckets
func TestInsert(t *testing.T) {
	bucket := InitBucket()
	
	err := Insert(5, bucket)
	if err != nil {
		t.Errorf("Insert should not return error: %v", err)
	}
	
	if len(bucket[5]) != 1 {
		t.Errorf("Expected bucket[5] to have 1 element, got %d", len(bucket[5]))
	}
	if bucket[5][0] != 5 {
		t.Errorf("Expected bucket[5][0] to be 5, got %d", bucket[5][0])
	}
}

// TestInsertOutOfBounds tests inserting out of bounds
func TestInsertOutOfBounds(t *testing.T) {
	bucket := InitBucket()
	
	err := Insert(B+1, bucket)
	if err == nil {
		t.Error("Insert should return error for out of bounds item")
	}
}

// TestPop tests popping items from buckets
func TestPop(t *testing.T) {
	bucket := InitBucket()
	
	_ = Insert(5, bucket)
	_ = Insert(2, bucket)
	_ = Insert(8, bucket)
	
	result := Pop(bucket)
	
	if len(result) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(result))
	}
	
	// Check order (should be ascending)
	expected := []int{2, 5, 8}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
		}
	}
}

// TestBucketSort tests the complete bucket sort
func TestBucketSort(t *testing.T) {
	items := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	
	err := BucketSort(items)
	if err != nil {
		t.Errorf("BucketSort should not return error: %v", err)
	}
	
	// Verify sorted
	for i := 0; i < len(items)-1; i++ {
		if items[i] > items[i+1] {
			t.Error("Array should be sorted in ascending order")
		}
	}
	
	// Verify all elements present
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range expected {
		if items[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, items[i])
		}
	}
}

// TestBucketSortDuplicates tests bucket sort with duplicates
func TestBucketSortDuplicates(t *testing.T) {
	items := []int{5, 2, 5, 1, 2}
	
	err := BucketSort(items)
	if err != nil {
		t.Errorf("BucketSort should not return error: %v", err)
	}
	
	expected := []int{1, 2, 2, 5, 5}
	for i, v := range expected {
		if items[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, items[i])
		}
	}
}

// TestBucketSortEmpty tests bucket sort with empty slice
func TestBucketSortEmpty(t *testing.T) {
	items := []int{}
	
	err := BucketSort(items)
	if err != nil {
		t.Errorf("BucketSort should not return error for empty slice: %v", err)
	}
	
	if len(items) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

// Bucket operations for atoms

type Átomo struct {
	Nombre string
	UMA    float64
}

// InitAtomBucket initializes a bucket for atoms
func InitAtomBucket() [][]Átomo {
	bucket := make([][]Átomo, 1+B)
	for i := 0; i <= B; i++ {
		bucket[i] = []Átomo{}
	}
	return bucket
}

// InsertAtom adds an atom to the bucket
func InsertAtom(átomo Átomo, bucket [][]Átomo) error {
	if int(átomo.UMA) >= len(bucket) {
		return fmt.Errorf("item out of bounds: %v > %v", átomo.UMA, B)
	}
	i := int(átomo.UMA)
	bucket[i] = append(bucket[i], átomo)
	return nil
}

// PopAtoms returns all atoms in ascending order by UMA
func PopAtoms(bucket [][]Átomo) []Átomo {
	result := []Átomo{}
	for i := 0; i <= B; i++ {
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)
		}
	}
	return result
}

// SortAtoms sorts atoms by UMA
func SortAtoms(items []Átomo) error {
	bckt := InitAtomBucket()
	
	for _, v := range items {
		if err := InsertAtom(v, bckt); err != nil {
			return err
		}
	}
	
	sorted := PopAtoms(bckt)
	copy(items, sorted)
	return nil
}

// TestAtomBucketSort tests sorting atoms
func TestAtomBucketSort(t *testing.T) {
	items := []Átomo{
		{"Be", 9.012},
		{"H", 1.008},
		{"Li", 6.967},
		{"He", 4.003},
	}
	
	err := SortAtoms(items)
	if err != nil {
		t.Errorf("SortAtoms should not return error: %v", err)
	}
	
	// Verify sorted by UMA
	expectedOrder := []string{"H", "He", "Li", "Be"}
	for i, name := range expectedOrder {
		if items[i].Nombre != name {
			t.Errorf("Expected %s at index %d, got %s", name, i, items[i].Nombre)
		}
	}
}

// Multimap operations

// index returns the index of item in slice, or -1 if not found
func index(item int, slice []int) int {
	for k, v := range slice {
		if v == item {
			return k
		}
	}
	return -1
}

// remove removes element at index from slice
func remove(idx int, slice []int) ([]int, error) {
	if idx < 0 || idx >= len(slice) {
		return slice, errors.New("removing out of bounds")
	}
	
	// Remove in O(1) by swapping with last element
	slice[idx] = slice[len(slice)-1]
	return slice[:len(slice)-1], nil
}

// Set inserts value in the container indexed by key
func Set(key string, value int, multimap map[string][]int) {
	if _, ok := multimap[key]; !ok {
		multimap[key] = []int{value}
	} else {
		multimap[key] = append(multimap[key], value)
	}
}

// Get returns all values indexed by key
func Get(key string, multimap map[string][]int) ([]int, bool) {
	if values, ok := multimap[key]; !ok {
		return []int{}, false
	} else {
		return values, true
	}
}

// Remove removes value from the container indexed by key
func Remove(key string, value int, multimap map[string][]int) bool {
	if values, ok := multimap[key]; ok {
		idx := index(value, values)
		found := idx >= 0
		if found {
			multimap[key], _ = remove(idx, multimap[key])
			if len(multimap[key]) == 0 {
				delete(multimap, key)
			}
			return true
		}
	}
	return false
}

// TestMultimapIndex tests finding index in slice
func TestMultimapIndex(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	
	idx := index(3, slice)
	if idx != 2 {
		t.Errorf("Expected index 2, got %d", idx)
	}
	
	idx = index(10, slice)
	if idx != -1 {
		t.Errorf("Expected index -1 for non-existent item, got %d", idx)
	}
}

// TestMultimapRemove tests removing elements from slice
func TestMultimapRemove(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	
	newSlice, err := remove(2, slice)
	if err != nil {
		t.Errorf("Remove should not return error: %v", err)
	}
	
	if len(newSlice) != 4 {
		t.Errorf("Expected length 4, got %d", len(newSlice))
	}
	
	// 3 should be removed (element at index 2 is swapped with last and removed)
	found := false
	for _, v := range newSlice {
		if v == 3 {
			found = true
		}
	}
	if found {
		t.Error("Element 3 should have been removed")
	}
}

// TestMultimapSet tests setting values in multimap
func TestMultimapSet(t *testing.T) {
	multimap := make(map[string][]int)
	
	Set("key1", 1, multimap)
	Set("key1", 2, multimap)
	Set("key2", 3, multimap)
	
	if len(multimap["key1"]) != 2 {
		t.Errorf("Expected 2 values for key1, got %d", len(multimap["key1"]))
	}
	if len(multimap["key2"]) != 1 {
		t.Errorf("Expected 1 value for key2, got %d", len(multimap["key2"]))
	}
}

// TestMultimapGet tests getting values from multimap
func TestMultimapGet(t *testing.T) {
	multimap := make(map[string][]int)
	Set("key1", 1, multimap)
	Set("key1", 2, multimap)
	
	values, ok := Get("key1", multimap)
	if !ok {
		t.Error("Key1 should exist")
	}
	if len(values) != 2 {
		t.Errorf("Expected 2 values, got %d", len(values))
	}
	
	_, ok = Get("nonexistent", multimap)
	if ok {
		t.Error("Nonexistent key should return false")
	}
}

// TestMultimapRemoveOperation tests removing values from multimap
func TestMultimapRemoveOperation(t *testing.T) {
	multimap := make(map[string][]int)
	Set("key1", 1, multimap)
	Set("key1", 2, multimap)
	Set("key1", 3, multimap)
	
	// Remove existing value
	removed := Remove("key1", 2, multimap)
	if !removed {
		t.Error("Should return true when value is removed")
	}
	
	values, _ := Get("key1", multimap)
	if len(values) != 2 {
		t.Errorf("Expected 2 values after removal, got %d", len(values))
	}
	
	// Remove non-existing value
	removed = Remove("key1", 10, multimap)
	if removed {
		t.Error("Should return false when value doesn't exist")
	}
	
	// Remove all values
	Remove("key1", 1, multimap)
	Remove("key1", 3, multimap)
	
	// Key should be deleted when no values remain
	_, ok := Get("key1", multimap)
	if ok {
		t.Error("Key should be deleted when all values are removed")
	}
}

// TestMultimapComplete tests complete multimap operations
func TestMultimapComplete(t *testing.T) {
	multimap := make(map[string][]int)
	
	// Add values
	Set("colors", 1, multimap)
	Set("colors", 2, multimap)
	Set("colors", 3, multimap)
	Set("numbers", 10, multimap)
	Set("numbers", 20, multimap)
	
	// Check values
	colors, ok := Get("colors", multimap)
	if !ok || len(colors) != 3 {
		t.Error("Colors should have 3 values")
	}
	
	numbers, ok := Get("numbers", multimap)
	if !ok || len(numbers) != 2 {
		t.Error("Numbers should have 2 values")
	}
	
	// Remove value
	Remove("colors", 2, multimap)
	colors, _ = Get("colors", multimap)
	if len(colors) != 2 {
		t.Error("Colors should have 2 values after removal")
	}
}
