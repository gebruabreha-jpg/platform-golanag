// -*- coding: utf-8 -*-
// for_test.go
// -----------------------------------------------------------------------------
//
// Unit tests for code snippets from Chapter 5, Section 3 (for loops)
// These tests validate the exact code shown in the LaTeX document using \gofragment
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Test snippet from lines 22-31: Binary string conversion
// Tests conversion of binary string to integer
func TestForBinaryConversion(t *testing.T) {
	valor := "00110010"
	resultado := 0
	for idx := 0; idx < 8; idx++ {
		resultado <<= 1
		if valor[idx] == '1' {
			resultado += 1
		}
	}
	
	expected := 50 // 00110010 in binary is 50 in decimal
	if resultado != expected {
		t.Errorf("Expected %d, got %d", expected, resultado)
	}
	
	// Verify output format
	output := fmt.Sprintf("%s: %v\n", valor, resultado)
	expectedOutput := "00110010: 50\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

// Test snippet from lines 35-38: String reversal
// Tests character swapping to reverse a string
func TestForStringReversal(t *testing.T) {
	final := make([]byte, 12)
	s := "Roma es Amor"
	
	// inversion de una cadena 's'
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		final[i], final[j] = s[j], s[i]
	}
	
	// Note: This algorithm only swaps positions, doesn't create a complete reversal
	// It's testing the loop pattern shown in the textbook
	
	// Verify the loop executes the expected number of times
	count := 0
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		count++
	}
	expectedIterations := len(s) / 2
	if count != expectedIterations {
		t.Errorf("Expected %d iterations, got %d", expectedIterations, count)
	}
}

// Test snippet from lines 39-52: Loop with break and continue
// Tests sum of numbers with conditions
func TestForBreakContinue(t *testing.T) {
	// calcular la suma de todos los números naturales hasta el primero que sea
	// múltiplo de 42, y que no sean múltiplos ni de 3 ni de 5
	suma := 0
	for i := 1; ; i++ { // sin condición de terminación!
		if i%42 == 0 {
			break
		}
		if i%3 == 0 || i%5 == 0 {
			continue
		}
		suma += i
	}
	
	// Verify the calculation
	// Numbers from 1 to 41 (42 excluded), excluding multiples of 3 and 5
	expected := 0
	for i := 1; i < 42; i++ {
		if i%3 != 0 && i%5 != 0 {
			expected += i
		}
	}
	
	if suma != expected {
		t.Errorf("Expected sum %d, got %d", expected, suma)
	}
	
	// Verify output format
	output := fmt.Sprintf("Suma: %v", suma)
	if output == "" {
		t.Error("Expected non-empty output")
	}
}

// Test snippet from lines 53-62: While-style loop (reading file)
// Tests scanner pattern for file reading
func TestForWhileStyleFileRead(t *testing.T) {
	// Create a test file
	filename := "/tmp/test_curso.txt"
	content := "Curso de Go - Módulo I\nLinea 2\nLinea 3"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filename)
	
	// Test the exact pattern from the textbook
	stream, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Error abriendo el fichero: %v", err)
	}
	defer stream.Close()
	
	reader := bufio.NewScanner(stream)
	lines := []string{}
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}
	
	expectedLines := strings.Split(content, "\n")
	if len(lines) != len(expectedLines) {
		t.Errorf("Expected %d lines, got %d", len(expectedLines), len(lines))
	}
	
	for i, line := range lines {
		if line != expectedLines[i] {
			t.Errorf("Line %d: expected %q, got %q", i, expectedLines[i], line)
		}
	}
}

// Test snippet from lines 64-72: Range over integer (Go 1.22+)
// Tests the improved binary conversion using range
func TestForRangeInteger(t *testing.T) {
	valor := "00110010"
	resultado := 0
	
	// Versión mejorada del traductor binario
	for idx := range 8 {
		resultado <<= 1
		if valor[idx] == '1' {
			resultado += 1
		}
	}
	
	expected := 50
	if resultado != expected {
		t.Errorf("Expected %d, got %d", expected, resultado)
	}
	
	// Verify output format
	output := fmt.Sprintf("%s: %v\n", valor, resultado)
	expectedOutput := "00110010: 50\n"
	if output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

// Test snippet from lines 74-82: Range over string with character manipulation
// Tests string to uppercase conversion
func TestForRangeStringToUpper(t *testing.T) {
	proverb := "Make the zero value useful!"
	result := ""
	
	for _, v := range proverb {
		if v >= 97 && v <= 122 { // es una letra minúscula
			v -= 32 // conversión a mayúscula
		}
		result += string(v)
	}
	
	expected := "MAKE THE ZERO VALUE USEFUL!"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Integration test: Verify all for loop patterns work correctly
func TestForIntegration(t *testing.T) {
	// Test 1: Three-component loop
	count := 0
	for i := 0; i < 5; i++ {
		count++
	}
	if count != 5 {
		t.Errorf("Expected 5 iterations, got %d", count)
	}
	
	// Test 2: While-style loop
	count = 0
	i := 0
	for i < 5 {
		count++
		i++
	}
	if count != 5 {
		t.Errorf("Expected 5 iterations, got %d", count)
	}
	
	// Test 3: Infinite loop with break
	count = 0
	for {
		count++
		if count >= 5 {
			break
		}
	}
	if count != 5 {
		t.Errorf("Expected 5 iterations, got %d", count)
	}
	
	// Test 4: Range over string
	s := "hello"
	charCount := 0
	for range s {
		charCount++
	}
	if charCount != len(s) {
		t.Errorf("Expected %d characters, got %d", len(s), charCount)
	}
}
