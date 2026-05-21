// -*- coding: utf-8 -*-
// if_test.go
// -----------------------------------------------------------------------------
//
// Unit tests for code snippets from Chapter 5, Section 2 (if statements)
// These tests validate the exact code shown in the LaTeX document using \gofragment
//

package main

import (
	"fmt"
	"os"
	"testing"
)

// Test snippet from lines 21-23: Basic if statement
// Tests the pattern of checking for zero value
func TestIfBasicCheck(t *testing.T) {
	n := 0
	errorOccurred := false
	
	if n == 0 { // en vez de "if !n"
		// log.Fatal would terminate the test, so we just flag it
		errorOccurred = true
	}
	
	if !errorOccurred {
		t.Error("Expected error condition to be triggered when n == 0")
	}
	
	// Test with non-zero value
	n = 1
	errorOccurred = false
	if n == 0 {
		errorOccurred = true
	}
	
	if errorOccurred {
		t.Error("Expected no error condition when n != 0")
	}
}

// Test snippet from lines 25-32: Leap year calculation
// Tests the leap year logic shown in the textbook
func TestIfLeapYear(t *testing.T) {
	testCases := []struct {
		año      int
		expected bool
	}{
		{2020, true},  // Divisible by 4, not by 100
		{2000, true},  // Divisible by 400
		{1900, false}, // Divisible by 100, not by 400
		{2019, false}, // Not divisible by 4
		{2024, true},  // Divisible by 4
	}
	
	for _, tc := range testCases {
		año := tc.año
		añoBisiesto := año%4 == 0 && (año%100 != 0 ||
			(año%100 == 0 && año%400 == 0))
		
		if añoBisiesto != tc.expected {
			t.Errorf("For year %d: expected %v, got %v", año, tc.expected, añoBisiesto)
		}
		
		// Verify the output logic (without actually printing)
		if añoBisiesto {
			result := fmt.Sprintf(" El año %v es bisiesto\n", año)
			if result == "" {
				t.Error("Expected output for leap year")
			}
		} else {
			result := fmt.Sprintf(" El año %v no es bisiesto\n", año)
			if result == "" {
				t.Error("Expected output for non-leap year")
			}
		}
	}
}

// Test snippet from lines 34-37: if with initialization
// Tests file writing with error handling
func TestIfWithInitialization(t *testing.T) {
	// Create a temporary file for testing
	tmpfile := "/tmp/test_curso.txt"
	defer os.Remove(tmpfile)
	
	// Test the exact pattern from the textbook
	contents := []byte("Curso de Go - Módulo I")
	errorOccurred := false
	
	if err := os.WriteFile(tmpfile, contents, 0644); err != nil {
		errorOccurred = true
	}
	
	if errorOccurred {
		t.Error("Expected successful file write")
	}
	
	// Verify file was written correctly
	readContents, err := os.ReadFile(tmpfile)
	if err != nil {
		t.Fatalf("Failed to read back written file: %v", err)
	}
	
	if string(readContents) != string(contents) {
		t.Errorf("File contents mismatch: expected %q, got %q", string(contents), string(readContents))
	}
	
	// Test error condition with invalid path
	errorOccurred = false
	if err := os.WriteFile("/invalid/path/file.txt", contents, 0644); err != nil {
		errorOccurred = true
	}
	
	if !errorOccurred {
		t.Error("Expected error when writing to invalid path")
	}
}

// Integration test: Verify the complete example compiles and runs
func TestIfCompleteExample(t *testing.T) {
	// Test that all if patterns work together
	n := 1
	if n == 0 {
		t.Error("Should not reach here with n=1")
	}
	
	año := 2020
	añoBisiesto := año%4 == 0 && (año%100 != 0 ||
		(año%100 == 0 && año%400 == 0))
	if !añoBisiesto {
		t.Error("2020 should be a leap year")
	}
	
	tmpfile := "/tmp/test_complete.txt"
	defer os.Remove(tmpfile)
	
	contents := []byte("Curso de Go - Módulo I")
	if err := os.WriteFile(tmpfile, contents, 0644); err != nil {
		t.Fatalf("File write failed: %v", err)
	}
}
