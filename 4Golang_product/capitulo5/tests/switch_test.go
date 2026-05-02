// -*- coding: utf-8 -*-
// switch_test.go
// -----------------------------------------------------------------------------
//
// Unit tests for code snippets from Chapter 5, Section 2 (switch statements)
// These tests validate the exact code shown in the LaTeX document using \gofragment
//

package main

import (
	"fmt"
	"testing"
)

// Test snippet from lines 16-24: Switch with first match only
// Demonstrates that only the first matching case executes
func TestSwitchFirstMatchOnly(t *testing.T) {
	val := 18
	result := ""
	
	switch {
	case val%2 == 0:
		result = "es múltiplo de 2"
	case val%3 == 0:
		result = "es múltiplo de 3"
	case val%6 == 0:
		result = "es múltiplo de 6"
	}
	
	// Even though 18 is divisible by 2, 3, and 6, only the first case executes
	if result != "es múltiplo de 2" {
		t.Errorf("Expected 'es múltiplo de 2', got '%s'", result)
	}
	
	// Test with a value divisible by 3 but not 2
	val = 9
	result = ""
	switch {
	case val%2 == 0:
		result = "es múltiplo de 2"
	case val%3 == 0:
		result = "es múltiplo de 3"
	case val%6 == 0:
		result = "es múltiplo de 6"
	}
	
	if result != "es múltiplo de 3" {
		t.Errorf("Expected 'es múltiplo de 3', got '%s'", result)
	}
}

// Test snippet from lines 26-37: Switch without variable (guard-based)
// Tests chess game state logic
func TestSwitchGuardBased(t *testing.T) {
	testCases := []struct {
		jaque       bool
		movimientos int
		expected    string
	}{
		{true, 0, "Jaque mate\n"},
		{false, 0, "Ahogado\n"},
		{true, 5, "Jaque\n"},
		{false, 42, "Su turno\n"},
	}
	
	for _, tc := range testCases {
		result := ""
		jaque := tc.jaque
		movimientos := tc.movimientos
		
		switch {
		case jaque && movimientos == 0:
			result = "Jaque mate\n"
		case !jaque && movimientos == 0:
			result = "Ahogado\n"
		case jaque && movimientos > 0:
			result = "Jaque\n"
		default:
			result = "Su turno\n"
		}
		
		if result != tc.expected {
			t.Errorf("For jaque=%v, movimientos=%d: expected %q, got %q",
				tc.jaque, tc.movimientos, tc.expected, result)
		}
	}
}

// Test snippet from lines 39-53: Switch with multiple values per case
// Tests month-to-days conversion with leap year logic
func TestSwitchMonthDays(t *testing.T) {
	testCases := []struct {
		mes          string
		añoBisiesto  bool
		expectedDias int
	}{
		{"Agosto", false, 31},
		{"Enero", false, 31},
		{"Abril", false, 30},
		{"Febrero", false, 28},
		{"Febrero", true, 29},
		{"Diciembre", false, 31},
	}
	
	for _, tc := range testCases {
		var dias int
		añoBisiesto := tc.añoBisiesto
		mes := tc.mes
		
		switch mes {
		case "Enero", "Marzo", "Mayo", "Julio", "Agosto", "Octubre", "Diciembre":
			dias = 31
		case "Abril", "Junio", "Septiembre", "Noviembre":
			dias = 30
		case "Febrero":
			dias = 28
			if añoBisiesto {
				dias = 29
			}
		}
		
		if dias != tc.expectedDias {
			t.Errorf("For %s (leap=%v): expected %d days, got %d",
				tc.mes, tc.añoBisiesto, tc.expectedDias, dias)
		}
		
		// Verify output formatting
		output := fmt.Sprintf("%v tiene %v días\n", mes, dias)
		if output == "" {
			t.Error("Expected non-empty output")
		}
	}
}

// Test snippet from lines 55-71: Switch with fallthrough
// Tests chess piece movement logic with fallthrough for Dama
func TestSwitchFallthrough(t *testing.T) {
	testCases := []struct {
		pieza           string
		expectedOutputs []string
	}{
		{"Peón", []string{"Calculando movimientos de peón"}},
		{"Caballo", []string{"Calculando movimientos en L"}},
		{"Alfil", []string{"Calculando movimientos diagonales"}},
		{"Dama", []string{"Calculando movimientos diagonales", "Calculando movimientos horizontales"}},
		{"Torre", []string{"Calculando movimientos horizontales"}},
		{"Rey", []string{"Calculando movimientos de un paso"}},
	}
	
	for _, tc := range testCases {
		outputs := []string{}
		pieza := tc.pieza
		
		switch pieza {
		case "Peón":
			outputs = append(outputs, "Calculando movimientos de peón")
		case "Caballo":
			outputs = append(outputs, "Calculando movimientos en L")
		case "Alfil", "Dama":
			outputs = append(outputs, "Calculando movimientos diagonales")
			if pieza == "Alfil" {
				break // fallthrough debe estar en el primer nivel
			}
			fallthrough
		case "Torre": // La Dama ejecuta el caso anterior y este
			outputs = append(outputs, "Calculando movimientos horizontales")
		case "Rey":
			outputs = append(outputs, "Calculando movimientos de un paso")
		}
		
		if len(outputs) != len(tc.expectedOutputs) {
			t.Errorf("For %s: expected %d outputs, got %d",
				tc.pieza, len(tc.expectedOutputs), len(outputs))
			continue
		}
		
		for i, expected := range tc.expectedOutputs {
			if outputs[i] != expected {
				t.Errorf("For %s, output %d: expected %q, got %q",
					tc.pieza, i, expected, outputs[i])
			}
		}
	}
}

// Integration test: Verify all switch patterns work correctly
func TestSwitchIntegration(t *testing.T) {
	// Test 1: First match
	val := 18
	matched := false
	switch {
	case val%2 == 0:
		matched = true
	case val%3 == 0:
		matched = true
	}
	if !matched {
		t.Error("Expected at least one case to match")
	}
	
	// Test 2: Guard-based
	jaque, movimientos := false, 42
	result := ""
	switch {
	case jaque && movimientos == 0:
		result = "Jaque mate"
	case !jaque && movimientos == 0:
		result = "Ahogado"
	default:
		result = "Su turno"
	}
	if result == "" {
		t.Error("Expected result to be set")
	}
	
	// Test 3: Multiple values
	mes := "Agosto"
	var dias int
	switch mes {
	case "Enero", "Marzo", "Mayo", "Julio", "Agosto", "Octubre", "Diciembre":
		dias = 31
	}
	if dias != 31 {
		t.Errorf("Expected 31 days for Agosto, got %d", dias)
	}
}
