// Test suite for Chapter 4: Strings (Sección 7)
package capitulo4_test

import (
"strings"
"testing"
"unicode/utf8"
)

// TestStringBasics tests basic string properties
func TestStringBasics(t *testing.T) {
str := "Errors are values"

if len(str) != 17 {
t.Errorf("Expected len=17, got %d", len(str))
}
}

// TestStringImmutability tests that strings are immutable
func TestStringImmutability(t *testing.T) {
// This test documents that strings are immutable
// str[0] = 'x' would be a compile error
str := "hello"

// We can only create new strings, not modify existing ones
newStr := "H" + str[1:]

if newStr != "Hello" {
t.Errorf("Expected 'Hello', got '%s'", newStr)
}

// Original string unchanged
if str != "hello" {
t.Errorf("Expected original string 'hello', got '%s'", str)
}
}

// TestStringIndexing tests string byte indexing
func TestStringIndexing(t *testing.T) {
str := "Go"

if str[0] != 'G' {
t.Errorf("Expected 'G', got '%c'", str[0])
}
if str[0] != 0x47 {
t.Errorf("Expected 0x47, got 0x%x", str[0])
}
}

// TestStringSlicing tests string slicing
func TestStringSlicing(t *testing.T) {
s := "G' Day"
result := s[3:] + s[:3]

if result != "DayG' " {
t.Errorf("Expected 'DayG' ', got '%s'", result)
}
}

// TestStringConcatenation tests string concatenation
func TestStringConcatenation(t *testing.T) {
s1 := "Hello"
s2 := "World"
result := s1 + " " + s2

if result != "Hello World" {
t.Errorf("Expected 'Hello World', got '%s'", result)
}
}

// TestStringLength tests len() vs utf8.RuneCountInString()
func TestStringLength(t *testing.T) {
formula := "sin(θ)=cos(π/2-θ)"

byteCount := len(formula)
runeCount := utf8.RuneCountInString(formula)

// The string has 20 bytes but 17 runes (θ and π take 2 bytes each)
if byteCount != 20 {
t.Errorf("Expected 20 bytes, got %d", byteCount)
}
if runeCount != 17 {
t.Errorf("Expected 17 runes, got %d", runeCount)
}
}

// TestUTF8Support tests UTF-8 string handling
func TestUTF8Support(t *testing.T) {
str := "Errors are values"

// ASCII characters are 1 byte each
if len(str) != utf8.RuneCountInString(str) {
t.Errorf("ASCII string should have equal byte and rune count")
}

// Greek letters take multiple bytes
greek := "θ"
if len(greek) != 2 {
t.Errorf("Expected θ to be 2 bytes, got %d", len(greek))
}
if utf8.RuneCountInString(greek) != 1 {
t.Errorf("Expected θ to be 1 rune, got %d", utf8.RuneCountInString(greek))
}
}

// TestRuneIteration tests iterating over runes with range
func TestRuneIteration(t *testing.T) {
formula := "sin(θ)"

runes := []rune{}
positions := []int{}

for pos, runa := range formula {
runes = append(runes, runa)
positions = append(positions, pos)
}

// Should have 6 runes: s, i, n, (, θ, )
if len(runes) != 6 {
t.Errorf("Expected 6 runes, got %d", len(runes))
}

// Position 4 should be θ (positions: 0,1,2,3,4,6 because θ takes 2 bytes)
if runes[4] != 'θ' {
t.Errorf("Expected rune at index 4 to be θ, got %c", runes[4])
}

// The byte position of θ should be 4
if positions[4] != 4 {
t.Errorf("Expected θ at byte position 4, got %d", positions[4])
}
}

// TestDecodeRune tests utf8.DecodeRuneInString
func TestDecodeRune(t *testing.T) {
str := "θ"

r, size := utf8.DecodeRuneInString(str)

if r != 'θ' {
t.Errorf("Expected to decode θ, got %c", r)
}
if size != 2 {
t.Errorf("Expected θ to be 2 bytes, got %d", size)
}
}

// TestRawStrings tests raw string literals
func TestRawStrings(t *testing.T) {
reg := `^\s*([A-Za-z0-9]+)\s*:\s*%([A-Za-z]+)\s*`

// Raw strings don't interpret escape sequences
if !strings.Contains(reg, `\s`) {
t.Error("Raw string should contain literal \\s")
}
}

// TestEscapeSequences tests escape sequences in regular strings
func TestEscapeSequences(t *testing.T) {
str := "Hello\nWorld"

if !strings.Contains(str, "\n") {
t.Error("String should contain newline")
}

// Test hex escape
hex := "\x47\x6f"  // "Go"
if hex != "Go" {
t.Errorf("Expected 'Go', got '%s'", hex)
}

// Test unicode escape
unicode := "\u2655"  // White chess queen
if len(unicode) != 3 {  // UTF-8 encoding of U+2655
t.Errorf("Expected 3 bytes for ♕, got %d", len(unicode))
}
}

// TestStringComparison tests string comparison
func TestStringComparison(t *testing.T) {
s1 := "hello"
s2 := "hello"
s3 := "world"

if s1 != s2 {
t.Error("Expected equal strings to compare equal")
}
if s1 == s3 {
t.Error("Expected different strings to compare not equal")
}
if s1 >= s3 {
t.Error("Expected 'hello' < 'world'")
}
}

// TestStringSubslicing tests creating substrings
func TestStringSubslicing(t *testing.T) {
str := "Errors are values"

errors := str[:6]
values := str[11:]

if errors != "Errors" {
t.Errorf("Expected 'Errors', got '%s'", errors)
}
if values != "values" {
t.Errorf("Expected 'values', got '%s'", values)
}
}

// TestEmptyString tests empty string handling
func TestEmptyString(t *testing.T) {
empty := ""

if len(empty) != 0 {
t.Error("Empty string should have length 0")
}
if utf8.RuneCountInString(empty) != 0 {
t.Error("Empty string should have 0 runes")
}
}

// TestByteSliceConversion tests conversion between []byte and string
func TestByteSliceConversion(t *testing.T) {
str := "hello"
bytes := []byte(str)

if len(bytes) != 5 {
t.Errorf("Expected 5 bytes, got %d", len(bytes))
}

newStr := string(bytes)
if newStr != str {
t.Errorf("Expected '%s', got '%s'", str, newStr)
}
}
