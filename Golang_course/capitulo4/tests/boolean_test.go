// Test suite for Chapter 4: Boolean Type (Sección 4)
package capitulo4_test

import (
"testing"
)

// TestBooleanConstants tests that true and false are the boolean constants
func TestBooleanConstants(t *testing.T) {
var b1 bool = true
var b2 bool = false

if b1 != true || b2 != false {
t.Error("Boolean constants not working as expected")
}
}

// TestBooleanNotRelatedToIntegers tests that booleans don't implicitly convert to 0 or 1
func TestBooleanNotRelatedToIntegers(t *testing.T) {
varBool := true
var varInt int

if varBool {
varInt = 1
} else {
varInt = 0
}

if varInt != 1 {
t.Error("Boolean to int conversion not working as expected")
}

varBool = false
if varBool {
varInt = 1
} else {
varInt = 0
}

if varInt != 0 {
t.Error("Boolean to int conversion not working as expected")
}
}

// TestBooleanNotNil tests that booleans are not related to nil
func TestBooleanNotNil(t *testing.T) {
var b bool

// Default value is false, not nil
if b != false {
t.Error("Boolean default should be false, not nil")
}
}

// TestUnaryNotOperator tests the unary ! (not) operator
func TestUnaryNotOperator(t *testing.T) {
eof := false
datos := !eof

if datos != true {
t.Errorf("Expected !false to be true, got %v", datos)
}

eof = true
datos = !eof

if datos != false {
t.Errorf("Expected !true to be false, got %v", datos)
}
}

// TestLogicalAndOperator tests the && (and) operator
func TestLogicalAndOperator(t *testing.T) {
tests := []struct {
a, b     bool
expected bool
}{
{true, true, true},
{true, false, false},
{false, true, false},
{false, false, false},
}

for _, tt := range tests {
result := tt.a && tt.b
if result != tt.expected {
t.Errorf("Expected %v && %v = %v, got %v", tt.a, tt.b, tt.expected, result)
}
}
}

// TestLogicalOrOperator tests the || (or) operator
func TestLogicalOrOperator(t *testing.T) {
tests := []struct {
a, b     bool
expected bool
}{
{true, true, true},
{true, false, true},
{false, true, true},
{false, false, false},
}

for _, tt := range tests {
result := tt.a || tt.b
if result != tt.expected {
t.Errorf("Expected %v || %v = %v, got %v", tt.a, tt.b, tt.expected, result)
}
}
}

// TestShortCircuitAnd tests that && is short-circuit
func TestShortCircuitAnd(t *testing.T) {
y := 10
x := 5

// This should not panic because of short-circuit
result := y > 0 && x/y > 0

if result != false {
t.Errorf("Expected false, got %v", result)
}

// This should not panic - left side is false, right side not evaluated
y = 0
result = y > 0 && x/y > 10  // x/y would panic if evaluated

if result != false {
t.Errorf("Expected false, got %v", result)
}
}

// TestShortCircuitOr tests that || is short-circuit
func TestShortCircuitOr(t *testing.T) {
s := "hello"

// This should not panic because of short-circuit
result := s != "" && s[0] < 'z'

if result != true {
t.Errorf("Expected true, got %v", result)
}

// Empty string test - should not panic
s = ""
result = s != "" && s[0] < 'z'  // s[0] would panic if s is empty and evaluated

if result != false {
t.Errorf("Expected false, got %v", result)
}
}

// TestBooleanExpressionExample tests the example from the chapter
func TestBooleanExpressionExample(t *testing.T) {
eof := false
nbytes := 1025

finlectura := eof || (nbytes > 1024)

if finlectura != true {
t.Errorf("Expected true, got %v", finlectura)
}
}
