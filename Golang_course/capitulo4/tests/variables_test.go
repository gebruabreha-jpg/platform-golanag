// Test suite for Chapter 4: Variables (Sección 3)
package capitulo4_test

import (
"testing"
)

// TestExplicitDeclaration tests explicit variable declarations
func TestExplicitDeclaration(t *testing.T) {
var edad int
var altura, peso float64
var nombre, apellidos string

// Check default values
if edad != 0 {
t.Errorf("Expected edad default to be 0, got %d", edad)
}
if altura != 0.0 || peso != 0.0 {
t.Error("Expected float default to be 0.0")
}
if nombre != "" || apellidos != "" {
t.Error("Expected string default to be empty string")
}
}

// TestFactorizedDeclaration tests factorized variable declarations
func TestFactorizedDeclaration(t *testing.T) {
var (
edad int
altura, peso float64
nombre, apellidos string
)

// Variables should be initialized with default values
if edad != 0 || altura != 0.0 || peso != 0.0 || nombre != "" || apellidos != "" {
t.Error("Factorized variables not initialized to default values")
}
}

// TestInitialization tests variable initialization
func TestInitialization(t *testing.T) {
var edad int = 9
var altura, peso float64 = 1.39, 28.2
var nombre, apellidos string = "Rob", "Pike"

if edad != 9 {
t.Errorf("Expected edad to be 9, got %d", edad)
}
if altura != 1.39 || peso != 28.2 {
t.Errorf("Expected altura=1.39, peso=28.2, got altura=%f, peso=%f", altura, peso)
}
if nombre != "Rob" || apellidos != "Pike" {
t.Errorf("Expected nombre=Rob, apellidos=Pike, got nombre=%s, apellidos=%s", nombre, apellidos)
}
}

// TestShortDeclaration tests short variable declaration and initialization
func TestShortDeclaration(t *testing.T) {
edad := 9
altura, peso := 1.31, 28.2
nombre, apellidos := "Rob", "Pike"

if edad != 9 {
t.Errorf("Expected edad to be 9, got %d", edad)
}
if altura != 1.31 || peso != 28.2 {
t.Errorf("Expected altura=1.31, peso=28.2, got altura=%f, peso=%f", altura, peso)
}
if nombre != "Rob" || apellidos != "Pike" {
t.Errorf("Expected nombre=Rob, apellidos=Pike, got nombre=%s, apellidos=%s", nombre, apellidos)
}
}

// TestVariableMutability tests that variables are mutable
func TestVariableMutability(t *testing.T) {
edad := 9
edad = 10

if edad != 10 {
t.Errorf("Expected edad to be mutable and changed to 10, got %d", edad)
}
}

// TestDefaultValues tests default values for all basic types
func TestDefaultValues(t *testing.T) {
var b bool
var i int
var f float64
var s string

if b != false {
t.Error("Expected bool default to be false")
}
if i != 0 {
t.Error("Expected int default to be 0")
}
if f != 0.0 {
t.Error("Expected float64 default to be 0.0")
}
if s != "" {
t.Error("Expected string default to be empty string")
}
}
