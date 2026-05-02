// Test suite for Chapter 4: Constants (Sección 2)
package capitulo4_test

import (
"math"
"testing"
)

// TestConstantsWithType tests explicit constant declarations with type
func TestConstantsWithType(t *testing.T) {
const base int = 2
const pi float64 = math.Pi

if base != 2 {
t.Errorf("Expected base to be 2, got %d", base)
}

if pi != math.Pi {
t.Errorf("Expected pi to be %f, got %f", math.Pi, pi)
}
}

// TestConstantsWithoutType tests implicit constant declarations without type
func TestConstantsWithoutType(t *testing.T) {
const chaitin = 0.0000001

if chaitin != 0.0000001 {
t.Errorf("Expected chaitin to be 0.0000001, got %f", chaitin)
}
}

// TestFactorizedConstants tests factorized constant declarations
func TestFactorizedConstants(t *testing.T) {
const (
base int = 2
pi float64 = math.Pi
chaitin = 0.0000001
)

if base != 2 || pi != math.Pi || chaitin != 0.0000001 {
t.Error("Factorized constants not correctly initialized")
}
}

// TestConstantReuse tests that constants reuse previous values in factorized blocks
func TestConstantReuse(t *testing.T) {
const (
peon = 1
caballo = 3
alfil          // should be 3 (reuses caballo)
torre = 5
dama = 9
)

if alfil != 3 {
t.Errorf("Expected alfil to be 3 (reusing caballo value), got %d", alfil)
}
}

// TestConstantWithUnderscore tests underscore as aesthetic marker
func TestConstantWithUnderscore(t *testing.T) {
const (
peon = 1
_              // aesthetic marker, doesn't force step
caballo = 3
alfil          // should be 3 (reuses caballo)
torre = 5
dama = 9
)

if alfil != 3 {
t.Errorf("Expected alfil to be 3, got %d", alfil)
}
}

// TestIotaBasic tests basic iota usage starting from 0
func TestIotaBasic(t *testing.T) {
const (
Lunes = iota
Martes
Miercoles
Jueves
Viernes
Sabado
Domingo
)

expected := map[int]string{
0: "Lunes",
1: "Martes",
2: "Miercoles",
3: "Jueves",
4: "Viernes",
5: "Sabado",
6: "Domingo",
}

if Lunes != 0 || Martes != 1 || Miercoles != 2 || Jueves != 3 ||
   Viernes != 4 || Sabado != 5 || Domingo != 6 {
t.Errorf("Iota sequence incorrect. Expected %v", expected)
}
}

// TestIotaWithExpression tests iota with expressions
func TestIotaWithExpression(t *testing.T) {
const (
enSusMarcas = 0b11 &^ iota  // 0b11 &^ 0 = 3
preparados                   // 0b11 &^ 1 = 2
listos                       // 0b11 &^ 2 = 1
ya                          // 0b11 &^ 3 = 0
)

if enSusMarcas != 3 || preparados != 2 || listos != 1 || ya != 0 {
t.Errorf("Expected 3,2,1,0 but got %d,%d,%d,%d", 
enSusMarcas, preparados, listos, ya)
}
}

// TestIotaStartingAt1 tests iota starting at 1 (like time.Month)
func TestIotaStartingAt1(t *testing.T) {
type Month int

const (
January Month = 1 + iota
February
March
April
May
June
July
August
September
October
November
December
)

if January != 1 || February != 2 || December != 12 {
t.Errorf("Month constants incorrect. January=%d, February=%d, December=%d",
January, February, December)
}
}
