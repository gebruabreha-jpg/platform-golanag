// Test suite for Chapter 4: Pointers (Sección 6)
package capitulo4_test

import (
"testing"
)

// TestPointerBasics tests basic pointer operations
func TestPointerBasics(t *testing.T) {
a := 5
var aref *int = &a

// Check that aref points to a
if *aref != 5 {
t.Errorf("Expected *aref to be 5, got %d", *aref)
}

// Modify through pointer
*aref = 10
if a != 10 {
t.Errorf("Expected a to be 10 after modifying through pointer, got %d", a)
}
}

// TestPointerReferencing tests the & (referencing) operator
func TestPointerReferencing(t *testing.T) {
var a int = 5
var aref *int = &a

// aref should hold the address of a
if aref == nil {
t.Error("Expected aref to hold address of a, got nil")
}

// Dereferencing should give original value
if *aref != 5 {
t.Errorf("Expected *aref to be 5, got %d", *aref)
}
}

// TestPointerDereferencing tests the * (dereferencing) operator
func TestPointerDereferencing(t *testing.T) {
var a int = 5
var aref *int = &a
var ad int = *aref

if ad != 5 {
t.Errorf("Expected ad to be 5, got %d", ad)
}
}

// TestPointerChaining tests chaining & and * operators
func TestPointerChaining(t *testing.T) {
var e int = 40
var eref *int = &e

// &*eref should equal eref
h := &*eref
if h != eref {
t.Error("Expected &*eref to equal eref")
}

// *&e should equal e
i := *&e
if i != e {
t.Error("Expected *&e to equal e")
}
}

// TestPointerDefaultValue tests that nil is the default value for pointers
func TestPointerDefaultValue(t *testing.T) {
var nilptr *int

if nilptr != nil {
t.Errorf("Expected nil pointer, got %v", nilptr)
}
}

// TestPointerTypes tests pointers to different types
func TestPointerTypes(t *testing.T) {
str := "hello"
strptr := &str

if *strptr != "hello" {
t.Errorf("Expected *strptr to be 'hello', got '%s'", *strptr)
}

f := 3.14
fptr := &f

if *fptr != 3.14 {
t.Errorf("Expected *fptr to be 3.14, got %f", *fptr)
}
}

// TestPointerModification tests modifying values through pointers
func TestPointerModification(t *testing.T) {
a := 5
aptr := &a

*aptr = 20

if a != 20 {
t.Errorf("Expected a to be 20 after modification through pointer, got %d", a)
}
}

// TestPointerComparison tests pointer comparison
func TestPointerComparison(t *testing.T) {
a := 5
b := 5

aptr := &a
aptr2 := &a
bptr := &b

// Two pointers to the same variable should be equal
if aptr != aptr2 {
t.Error("Expected pointers to same variable to be equal")
}

// Pointers to different variables should not be equal
if aptr == bptr {
t.Error("Expected pointers to different variables to be different")
}
}

// TestPointerToPointer tests pointer to pointer
func TestPointerToPointer(t *testing.T) {
a := 5
aptr := &a
aptrptr := &aptr

if **aptrptr != 5 {
t.Errorf("Expected **aptrptr to be 5, got %d", **aptrptr)
}
}
