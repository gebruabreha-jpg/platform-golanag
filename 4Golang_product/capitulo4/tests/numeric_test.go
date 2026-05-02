// Test suite for Chapter 4: Numeric Types (Sección 5)
package capitulo4_test

import (
"math"
"testing"
)

// TestIntegerTypes tests various integer type declarations
func TestIntegerTypes(t *testing.T) {
var i8 int8 = 127
var i16 int16 = 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807
var i int = 100

if i8 != 127 || i16 != 32767 || i32 != 2147483647 || i64 != 9223372036854775807 || i != 100 {
t.Error("Integer types not working as expected")
}
}

// TestUnsignedIntegerTypes tests unsigned integer types
func TestUnsignedIntegerTypes(t *testing.T) {
var ui8 uint8 = 255
var ui16 uint16 = 65535
var ui32 uint32 = 4294967295
var ui64 uint64 = 18446744073709551615
var ui uint = 100

if ui8 != 255 || ui16 != 65535 || ui32 != 4294967295 || ui64 != 18446744073709551615 || ui != 100 {
t.Error("Unsigned integer types not working as expected")
}
}

// TestIntegerAliases tests byte and rune aliases
func TestIntegerAliases(t *testing.T) {
var b byte = 255  // alias for uint8
var r rune = 'θ'  // alias for int32

if b != 255 {
t.Errorf("Expected byte to be 255, got %d", b)
}
if r != 'θ' {
t.Errorf("Expected rune to be θ, got %c", r)
}
}

// TestMathConstants tests math package constants for integers
func TestMathConstants(t *testing.T) {
if math.MaxInt8 != 127 {
t.Errorf("Expected MaxInt8=127, got %d", math.MaxInt8)
}
if math.MinInt8 != -128 {
t.Errorf("Expected MinInt8=-128, got %d", math.MinInt8)
}
if math.MaxUint8 != 255 {
t.Errorf("Expected MaxUint8=255, got %d", math.MaxUint8)
}
}

// TestUnaryIntegerOperators tests unary operators
func TestUnaryIntegerOperators(t *testing.T) {
a := 5

if -a != -5 {
t.Errorf("Expected -5, got %d", -a)
}
if +a != 5 {
t.Errorf("Expected 5, got %d", +a)
}
if ^a != -6 {  // bitwise NOT: ^5 = -6 in two's complement
t.Errorf("Expected -6, got %d", ^a)
}
}

// TestArithmeticOperators tests arithmetic operators
func TestArithmeticOperators(t *testing.T) {
a, b := 10, 3

if a+b != 13 {
t.Errorf("Expected 10+3=13, got %d", a+b)
}
if a-b != 7 {
t.Errorf("Expected 10-3=7, got %d", a-b)
}
if a*b != 30 {
t.Errorf("Expected 10*3=30, got %d", a*b)
}
if a/b != 3 {
t.Errorf("Expected 10/3=3, got %d", a/b)
}
if a%b != 1 {
t.Errorf("Expected 10%%3=1, got %d", a%b)
}
}

// TestModuloOperator tests modulo with positive and negative numbers
func TestModuloOperator(t *testing.T) {
// Positive dividend
if 34%13 != 8 {
t.Errorf("Expected 34%%13=8, got %d", 34%13)
}

// Negative dividend
if -34%13 != -8 {
t.Errorf("Expected -34%%13=-8, got %d", -34%13)
}
}

// TestBitwiseOperators tests bitwise operators
func TestBitwiseOperators(t *testing.T) {
a, b := 0b1100, 0b1010  // 12, 10

if a&b != 0b1000 {  // 8
t.Errorf("Expected 12&10=8, got %d", a&b)
}
if a|b != 0b1110 {  // 14
t.Errorf("Expected 12|10=14, got %d", a|b)
}
if a^b != 0b0110 {  // 6
t.Errorf("Expected 12^10=6, got %d", a^b)
}
if a&^b != 0b0100 {  // 4 (AND NOT)
t.Errorf("Expected 12&^10=4, got %d", a&^b)
}
}

// TestShiftOperators tests shift operators
func TestShiftOperators(t *testing.T) {
temperatura := -34
temperatura <<= 2

if temperatura != -136 {
t.Errorf("Expected -34<<2=-136, got %d", temperatura)
}

a := 16
if a>>2 != 4 {
t.Errorf("Expected 16>>2=4, got %d", a>>2)
}
}

// TestFloatTypes tests floating-point types
func TestFloatTypes(t *testing.T) {
var f32 float32 = 3.14
var f64 float64 = 3.14159265359

if f32 < 3.13 || f32 > 3.15 {
t.Errorf("Expected f32 ~3.14, got %f", f32)
}
if f64 < 3.14 || f64 > 3.15 {
t.Errorf("Expected f64 ~3.14159, got %f", f64)
}
}

// TestFloatFormats tests different float literal formats
func TestFloatFormats(t *testing.T) {
altura := 1.88
peso := 92.3
tierra := 5.972e+24
neptuno := 1.024e+26
h := 53e-12

if altura != 1.88 || peso != 92.3 {
t.Error("Float literals not working")
}
if tierra <= 0 || neptuno <= 0 || h <= 0 {
t.Error("Scientific notation not working")
}
}

// TestFloatOperators tests floating-point operators
func TestFloatOperators(t *testing.T) {
altura, peso := 1.88, 92.3
imc := peso / (altura * altura)

expected := 92.3 / (1.88 * 1.88)
if imc < expected-0.01 || imc > expected+0.01 {
t.Errorf("Expected IMC ~%f, got %f", expected, imc)
}
}

// TestMathMod tests math.Mod function
func TestMathMod(t *testing.T) {
peso := 92.3
altura := 1.88
resto := math.Mod(peso, altura)

// Expected: 0.18 (approximately)
if resto < 0.17 || resto > 0.19 {
t.Errorf("Expected resto ~0.18, got %f", resto)
}
}

// TestNaN tests NaN handling
func TestNaN(t *testing.T) {
nan := math.NaN()

// NaN is not equal to itself
if nan == nan {
t.Error("NaN should not equal itself")
}

// Use IsNaN to check
if !math.IsNaN(nan) {
t.Error("Expected IsNaN to return true")
}
}

// TestComplexNumbers tests complex number types
func TestComplexNumbers(t *testing.T) {
c1 := complex(1, 2)
c2 := complex(2, -3)

if real(c1) != 1 || imag(c1) != 2 {
t.Errorf("Expected c1=1+2i, got %v", c1)
}
if real(c2) != 2 || imag(c2) != -3 {
t.Errorf("Expected c2=2-3i, got %v", c2)
}
}

// TestComplexLiterals tests complex literal syntax
func TestComplexLiterals(t *testing.T) {
c := 1 - 2i

if real(c) != 1 || imag(c) != -2 {
t.Errorf("Expected c=1-2i, got %v", c)
}
}

// TestComplexOperators tests complex operators
func TestComplexOperators(t *testing.T) {
c1 := complex(1, 2)
c2 := complex(2, -3)
sum := c1 + c2

if real(sum) != 3 || imag(sum) != -1 {
t.Errorf("Expected sum=3-1i, got %v", sum)
}

neg := -complex(1, -2)
if real(neg) != -1 || imag(neg) != 2 {
t.Errorf("Expected neg=-1+2i, got %v", neg)
}
}

// TestMinMaxBuiltins tests min and max built-in functions (Go 1.21+)
func TestMinMaxBuiltins(t *testing.T) {
a, b := 5, 10

if min(a, b) != 5 {
t.Errorf("Expected min(5,10)=5, got %d", min(a, b))
}
if max(a, b) != 10 {
t.Errorf("Expected max(5,10)=10, got %d", max(a, b))
}
}

// TestMinMaxVariadic tests min/max with multiple arguments
func TestMinMaxVariadic(t *testing.T) {
x, y, z := 3.14, 2.71, 1.41

mínimo := min(x, y, z)
máximo := max(x, y, z)

if mínimo != 1.41 {
t.Errorf("Expected min=1.41, got %f", mínimo)
}
if máximo != 3.14 {
t.Errorf("Expected max=3.14, got %f", máximo)
}
}
