# Chapter 6 Tests - Composite Data Types

This directory contains comprehensive unit tests for all code examples shown in Chapter 6 (Tipos de datos compuestos) of the Go course slides.

## Overview

The tests cover the following topics from Chapter 6:

1. **Arrays** - Static indexed sequences
2. **Slices** - Dynamic indexed sequences  
3. **Maps** - Dynamic associative containers
4. **Structs** - Aggregate types
5. **Bucket Sort** - O(n) sorting algorithm implementation
6. **Multimaps** - Maps with multiple values per key

## Test Files

### arrays_test.go
Tests all array-related concepts and examples:
- Array declaration with const size
- Array comparison and type signatures
- Array operators (random access and slicing)
- Various initialization methods (literal, partial, auto-size with `...`, selective, indexed)
- Array mutability
- Pass by value vs. pass by reference
- Array to slice conversion with `[:]`
- Default values
- Fibonacci sequence example

**Coverage:** 11 test functions covering all array examples from seccion2.tex

### slices_test.go
Tests all slice-related concepts and examples:
- Slice declaration and nil behavior
- Slice structure (pointer, len, cap)
- Slice comparison function implementation
- Random access and slicing operators
- Array to slice conversion
- Rune slice conversion from strings
- Creating slices with `make()`
- Capacity and length behavior
- Default values and initialization
- The `append()` function (non-destructive)
- Pass by value vs. pass by reference
- The `copy()` function
- The `clear()` function (Go 1.21+)
- Literal initialization and sub-slicing

**Coverage:** 22 test functions covering all slice examples from seccion3.tex

### maps_test.go
Tests all map-related concepts and examples:
- Map declaration and nil behavior
- Map creation with `make()` and literal syntax
- Map initialization with literals
- Read and write operations with `[]` operator
- Default values for missing keys
- Checking key existence with the comma-ok idiom
- Deleting entries with `delete()`
- The `clear()` function (Go 1.21+)
- Map iteration with `range`
- Various key and value types (int, string, struct, slice, nested maps)
- Capacity hints
- Zero value behavior

**Coverage:** 14 test functions covering all map examples from seccion4.tex

### structs_test.go
Tests all struct-related concepts and examples:
- Struct declaration and default values
- Literal initialization (with and without field names)
- Partial initialization
- Struct comparison
- Field access (read and write)
- Structs containing slices
- Positional initialization
- Anonymous fields and struct embedding
- Field shadowing in embedded structs
- Pass by value vs. pass by reference
- Self-referential structs (linked lists)
- Nested structs
- Slices of structs

**Coverage:** 16 test functions covering all struct examples from seccion5.tex

### bucket_multimap_test.go
Tests the practical examples shown in the slides:

#### Bucket Sort (O(n) sorting)
- Bucket initialization
- Inserting items into buckets
- Out of bounds error handling
- Popping items in sorted order
- Complete bucket sort implementation
- Sorting with duplicates
- Empty slice handling
- Atom sorting by atomic mass (UMA)

**Coverage:** 8 test functions for integer bucket sort + 1 for atom sorting

#### Multimap Implementation
- Finding index in slice
- Removing elements from slice (O(1) removal)
- Setting values in multimap
- Getting values from multimap
- Removing values from multimap
- Automatic key deletion when empty
- Complete multimap workflow

**Coverage:** 7 test functions covering all multimap operations

## Running the Tests

### Run All Tests
```bash
cd code/capitulo6/tests
go test
```

### Run Tests with Verbose Output
```bash
go test -v
```

### Run Specific Test File
```bash
go test -v -run TestArray
```

### Run a Specific Test
```bash
go test -v -run TestArrayDeclaration
```

### Run Tests with Coverage Report
```bash
go test -cover
```

### Generate Detailed Coverage Report
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Run Tests Multiple Times (for race conditions)
```bash
go test -count=10
```

## Test Coverage Summary

| File | Test Functions | Coverage |
|------|----------------|----------|
| arrays_test.go | 11 | All array examples from seccion2.tex |
| slices_test.go | 22 | All slice examples from seccion3.tex |
| maps_test.go | 14 | All map examples from seccion4.tex |
| structs_test.go | 16 | All struct examples from seccion5.tex |
| bucket_multimap_test.go | 16 | All bucket sort and multimap examples |
| **Total** | **79** | **Complete coverage of Chapter 6** |

## Detailed Test Coverage

### Arrays (seccion2.tex)
✅ Array declaration with const size  
✅ Array size validation at compile time  
✅ Array comparison (same and different types)  
✅ Array type signature (size is part of type)  
✅ Random access operator `[i]`  
✅ Slicing operator `[i:j]`  
✅ Literal initialization  
✅ Partial initialization with default values  
✅ Auto-size with `...`  
✅ Selective initialization with indices  
✅ Indexed initialization with `iota`  
✅ Array mutability (unlike strings)  
✅ Pass by value (no modification)  
✅ Pass by reference (allows modification)  
✅ Array to slice conversion with `[:]`  
✅ Fibonacci sequence example  

### Slices (seccion3.tex)
✅ Slice declaration and initialization  
✅ Nil slice vs. empty slice  
✅ Slice structure (ptr, len, cap)  
✅ Slice comparison function (slices not directly comparable)  
✅ Random access with `[]`  
✅ Slicing operator `[:]` on arrays  
✅ String to rune slice conversion  
✅ Creating slices with `make(len[, cap])`  
✅ Panic when exceeding length  
✅ Cannot access beyond length even with capacity  
✅ Default values by type  
✅ `append()` function (returns new slice)  
✅ Modifying slice elements in-place (pass by value)  
✅ Append doesn't work with pass by value (modifies ptr/len/cap)  
✅ Pass by reference solution for append  
✅ `copy()` function (copies min of src and dst lengths)  
✅ `clear()` function (Go 1.21+)  
✅ Slicing to create sub-slices  
✅ Checking for nil vs. empty slices  
✅ Type distinction between slices and arrays  

### Maps (seccion4.tex)
✅ Map declaration and nil behavior  
✅ Map creation with `make(map[k]T)`  
✅ Map literal syntax `map[k]T{}`  
✅ Literal initialization with data  
✅ Read and write with `[]` operator  
✅ Default value for missing keys  
✅ Comma-ok idiom for key existence  
✅ `delete()` function  
✅ Deleting non-existent keys (no error)  
✅ `clear()` function (Go 1.21+)  
✅ Iteration with `range`  
✅ Various key types (comparable types)  
✅ Various value types (including slices, maps)  
✅ Capacity hint in `make()`  
✅ Zero value behavior  

### Structs (seccion5.tex)
✅ Struct declaration  
✅ Default values (zero values of field types)  
✅ Literal initialization with field names  
✅ Partial initialization  
✅ Struct comparison  
✅ Field access (read/write)  
✅ Structs containing slices  
✅ Positional initialization (without field names)  
✅ Anonymous fields (embedding)  
✅ Direct access to embedded fields  
✅ Field shadowing in embedded structs  
✅ Pass by value (no modification)  
✅ Pass by reference (allows modification)  
✅ Self-referential structs (pointer to same type)  
✅ Nested structs  
✅ Slices of structs  

### Bucket Sort Example
✅ Bucket initialization with B+1 entries  
✅ Inserting items into buckets  
✅ Out of bounds error handling  
✅ Popping items in ascending order  
✅ O(n) sorting implementation  
✅ Handling duplicates  
✅ Empty slice handling  
✅ Sorting structs (atoms by UMA)  

### Multimap Example
✅ Finding element index in slice  
✅ O(1) removal (swap with last, remove last)  
✅ Setting values (create or append)  
✅ Getting values with existence check  
✅ Removing specific values  
✅ Automatic key deletion when empty  
✅ Complete multimap operations  

## Go Version Compatibility

These tests are compatible with Go 1.21+ due to the use of:
- `clear()` built-in function (introduced in Go 1.21)

For earlier Go versions, the `clear()` tests can be commented out or removed.

## Notes on Test Implementation

### Random Seed (Deprecated in Go 1.20+)
The slides originally used `rand.Seed(time.Now().UnixNano())`, which is deprecated in Go 1.20+. The LaTeX has been updated to note that Go 1.20+ automatically seeds the random number generator.

### Panic Testing
Several tests verify that panics occur when expected (e.g., accessing beyond slice length). These tests use `defer` and `recover()` to catch panics.

### Order Independence
Map iteration order is not guaranteed in Go, so tests that iterate over maps are designed to be order-independent.

### Slice Modification
The tests demonstrate the important distinction between modifying slice elements (works with pass-by-value) and modifying slice structure like len/cap (requires pass-by-reference).

## Additional Resources

- [Go Specification - Array types](https://go.dev/ref/spec#Array_types)
- [Go Specification - Slice types](https://go.dev/ref/spec#Slice_types)
- [Go Specification - Map types](https://go.dev/ref/spec#Map_types)
- [Go Specification - Struct types](https://go.dev/ref/spec#Struct_types)
- [Go Blog - Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Go Blog - Go maps in action](https://go.dev/blog/maps)

## License

Part of the intro-go course materials from Universidad Carlos III de Madrid.
