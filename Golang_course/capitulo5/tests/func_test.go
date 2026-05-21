// -*- coding: utf-8 -*-
// func_test.go
// -----------------------------------------------------------------------------
//
// Unit tests for code snippets from Chapter 5, Section 4 (functions)
// These tests validate the exact code shown in the LaTeX document using \gofragment
//

package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os/exec"
	"strings"
	"testing"
)

// Test snippet from lines 23-29: Function with named return value
// Tests division function with early return for NaN
func TestFuncDiv1(t *testing.T) {
	// Define the function exactly as shown in the textbook
	div1 := func(a, b float64) (result float64) {
		if b == 0 {
			return math.NaN()
		}
		result = a / b
		return
	}
	
	// Test normal division
	result := div1(13, 4)
	expected := 3.25
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
	
	// Test division by zero
	result = div1(13, 0)
	if !math.IsNaN(result) {
		t.Errorf("Expected NaN, got %f", result)
	}
}

// Test snippet from lines 31-36: Multi-valued function
// Tests division returning both result and error
func TestFuncDiv2(t *testing.T) {
	div2 := func(x, y float64) (float64, error) {
		if y == 0 {
			return math.NaN(), errors.New("Overflow")
		}
		return x / y, nil
	}
	
	// Test normal division
	result, err := div2(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Expected 5.0, got %f", result)
	}
	
	// Test division by zero
	result, err = div2(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero")
	}
	if !math.IsNaN(result) {
		t.Errorf("Expected NaN, got %f", result)
	}
}

// Test snippet from lines 38-50: Fibonacci function
// Tests multi-valued function with overflow detection
func TestFuncFibonacci(t *testing.T) {
	// Devuelve los términos (i-1)- e i-ésimo de la serie
	// de Fibonacci respectivamente, a partir de Fib(0)=x
	// y Fib(1)=y, para i>1
	fibonacci := func(x, y, i uint64) (uint64, uint64, error) {
		n1, n0, n := y, x+y, uint64(2)
		for n < i {
			n1, n0, n = n0, n1+n0, n+1
			if n0 < n1 {
				return 0, 0, errors.New("Overflow")
			}
		}
		return n1, n0, nil
	}
	
	// Test F(2) with F(0)=0, F(1)=1
	_, result, err := fibonacci(0, 1, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1 {
		t.Errorf("Expected F(2)=1, got %d", result)
	}
	
	// Test F(6)
	_, result, err = fibonacci(0, 1, 6)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 8 {
		t.Errorf("Expected F(6)=8, got %d", result)
	}
	
	// Test F(10)
	_, result, err = fibonacci(0, 1, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 55 {
		t.Errorf("Expected F(10)=55, got %d", result)
	}
}

// Test snippet from lines 52-60: Function consuming multi-valued result
// Tests golden ratio calculation using Fibonacci
func TestFuncPhi(t *testing.T) {
	fibonacci := func(x, y, i uint64) (uint64, uint64, error) {
		n1, n0, n := y, x+y, uint64(2)
		for n < i {
			n1, n0, n = n0, n1+n0, n+1
			if n0 < n1 {
				return 0, 0, errors.New("Overflow")
			}
		}
		return n1, n0, nil
	}
	
	// Devuelve una estimación del valor de la razón áurea como el
	// cociente entre dos términos consecutivos de la serie de
	// Fibonacci y una indicación de si son válidos
	phi := func(beforelast, last uint64, err error) (float64, error) {
		if err != nil {
			return math.NaN(), err
		}
		return float64(last) / float64(beforelast), nil
	}
	
	// Test with F(50)
	razónÁurea, err := phi(fibonacci(0, 1, 50))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	// Golden ratio is approximately 1.618
	expected := 1.618033988749895
	if math.Abs(razónÁurea-expected) > 0.000000000000001 {
		t.Errorf("Expected ~%f, got %f", expected, razónÁurea)
	}
}

// Test snippet from lines 62-75: Function with external command
// Tests invoking sort command
func TestFuncOrdenar(t *testing.T) {
	// invoca el comando sort sobre un conjunto de líneas dadas como
	// un string con '\n'. Devuelve otra cadena con las líneas ordenadas
	// y una indicación de error si lo hubiera
	ordenar := func(lineas string) (string, error) {
		cmd := exec.Command("sort")
		cmd.Stdin = strings.NewReader(lineas)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("Error durante la ejecución de 'sort': %v", err)
		}
		return out.String(), nil
	}
	
	input := "Wilhelm Leibniz\nAlan Turing\nGottlob Frege"
	result, err := ordenar(input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	
	expected := "Alan Turing\nGottlob Frege\nWilhelm Leibniz\n"
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Test snippet from lines 77-80: Pass by value (swap)
// Tests that swap doesn't work with pass by value
func TestFuncSwapByValue(t *testing.T) {
	// intercambio (fallido) de los contenidos de dos variables enteras
	swap := func(a, b int) {
		a, b = b, a
	}
	
	op1, op2 := 2, 3
	swap(op1, op2)
	
	// Values should NOT change (pass by value)
	if op1 != 2 || op2 != 3 {
		t.Errorf("Expected op1=2, op2=3, got op1=%d, op2=%d", op1, op2)
	}
}

// Test snippet from lines 108-109: Function calls in main
// Tests calling div1 function
func TestFuncDiv1Calls(t *testing.T) {
	div1 := func(a, b float64) (result float64) {
		if b == 0 {
			return math.NaN()
		}
		result = a / b
		return
	}
	
	result1 := div1(13, 4)
	if result1 != 3.25 {
		t.Errorf("Expected 3.25, got %f", result1)
	}
	
	result2 := div1(13, 0)
	if !math.IsNaN(result2) {
		t.Errorf("Expected NaN, got %f", result2)
	}
}

// Test snippet from lines 112-117: Variable shadowing
// Tests scope and shadowing behavior
func TestFuncShadowing(t *testing.T) {
	const x = "interface{} says nothing"
	results := []int{}
	
	for x := range 3 {
		x := x * 2 // oculta el contador del bucle
		results = append(results, x)
	}
	
	// Inner x shadows loop counter, so we get [0, 2, 4]
	expected := []int{0, 2, 4}
	if len(results) != len(expected) {
		t.Errorf("Expected %d results, got %d", len(expected), len(results))
	}
	
	for i, v := range results {
		if v != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], v)
		}
	}
	
	// Outer x should still be the constant
	if x != "interface{} says nothing" {
		t.Errorf("Expected constant x unchanged, got %q", x)
	}
}

// Test snippet from lines 120-125: Multi-valued function usage
// Tests using div2 with error handling
func TestFuncDiv2Usage(t *testing.T) {
	div2 := func(x, y float64) (float64, error) {
		if y == 0 {
			return math.NaN(), errors.New("Overflow")
		}
		return x / y, nil
	}
	
	output := ""
	if res, err := div2(10, 0); err == nil {
		output = fmt.Sprintf("%v\n", res)
	} else {
		output = fmt.Sprintf("%v\n", err)
	}
	
	if output != "Overflow\n" {
		t.Errorf("Expected 'Overflow\\n', got %q", output)
	}
}

// Test snippet from lines 128-132: Composing multi-valued functions
// Tests calling phi with fibonacci result
func TestFuncPhiComposition(t *testing.T) {
	fibonacci := func(x, y, i uint64) (uint64, uint64, error) {
		n1, n0, n := y, x+y, uint64(2)
		for n < i {
			n1, n0, n = n0, n1+n0, n+1
			if n0 < n1 {
				return 0, 0, errors.New("Overflow")
			}
		}
		return n1, n0, nil
	}
	
	phi := func(beforelast, last uint64, err error) (float64, error) {
		if err != nil {
			return math.NaN(), err
		}
		return float64(last) / float64(beforelast), nil
	}
	
	if razónÁurea, err := phi(fibonacci(0, 1, 50)); err == nil {
		expected := 1.618033988749895
		if math.Abs(razónÁurea-expected) > 0.000000000000001 {
			t.Errorf("Expected ~%f, got %f", expected, razónÁurea)
		}
	} else {
		t.Errorf("Unexpected error: %v", err)
	}
}

// Test snippet from lines 134-139: Using ordenar function
// Tests complete sort example
func TestFuncOrdenarUsage(t *testing.T) {
	ordenar := func(lineas string) (string, error) {
		cmd := exec.Command("sort")
		cmd.Stdin = strings.NewReader(lineas)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("Error durante la ejecución de 'sort': %v", err)
		}
		return out.String(), nil
	}
	
	var result string
	var err error
	if result, err = ordenar("Wilhelm Leibniz\nAlan Turing\nGottlob Frege"); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	
	expected := "Alan Turing\nGottlob Frege\nWilhelm Leibniz\n"
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

// Test snippet from lines 142-145: Pass by value demonstration
// Tests that swap doesn't modify original values
func TestFuncSwapDemo(t *testing.T) {
	swap := func(a, b int) {
		a, b = b, a
	}
	
	op1, op2 := 2, 3
	before1, before2 := op1, op2
	
	swap(op1, op2)
	
	// Values should be unchanged (pass by value)
	if op1 != before1 || op2 != before2 {
		t.Errorf("Expected values unchanged: op1=%d, op2=%d, but got op1=%d, op2=%d",
			before1, before2, op1, op2)
	}
}

// Integration test: Verify all function patterns work together
func TestFuncIntegration(t *testing.T) {
	// Test function with named return
	div1 := func(a, b float64) (result float64) {
		if b == 0 {
			return math.NaN()
		}
		result = a / b
		return
	}
	
	if result := div1(10, 2); result != 5.0 {
		t.Errorf("div1: expected 5.0, got %f", result)
	}
	
	// Test multi-valued function
	div2 := func(x, y float64) (float64, error) {
		if y == 0 {
			return math.NaN(), errors.New("Overflow")
		}
		return x / y, nil
	}
	
	if res, err := div2(10, 2); err != nil || res != 5.0 {
		t.Errorf("div2: expected 5.0 with no error, got %f, %v", res, err)
	}
	
	// Test pass by value
	swap := func(a, b int) {
		a, b = b, a
	}
	
	x, y := 1, 2
	swap(x, y)
	if x != 1 || y != 2 {
		t.Error("swap should not modify original values")
	}
}
