package tests

import (
	"testing"
	"time"
)

// Planeta represents a planet
type Planeta struct {
	Símbolo            rune
	DiametroEcuatorial int
	PeríodoOrbital     float64
	Inclinación        float64
}

// Point represents a 3D point
type Point struct {
	X, Y, Z int
}

// TimedPosition represents a position with timestamp
type TimedPosition struct {
	Point
	T time.Time
}

// PlanetaExtended extends Planeta with position
type PlanetaExtended struct {
	Planeta
	TimedPosition
}

// Galaxia represents a galaxy
type Galaxia struct {
	Nombre    string
	Velocidad int
	Planetas  []Planeta
}

// TestStructDeclaration tests basic struct declarations
func TestStructDeclaration(t *testing.T) {
	var p Planeta
	
	// Check default values
	if p.Símbolo != 0 {
		t.Error("Default value for rune should be 0")
	}
	if p.DiametroEcuatorial != 0 {
		t.Error("Default value for int should be 0")
	}
	if p.PeríodoOrbital != 0.0 {
		t.Error("Default value for float64 should be 0.0")
	}
}

// TestStructLiteralInitialization tests struct initialization with literals
func TestStructLiteralInitialization(t *testing.T) {
	mercury := Planeta{
		Símbolo:            '☿',
		DiametroEcuatorial: 4878,
		PeríodoOrbital:     0.24,
		Inclinación:        7.0,
	}
	
	if mercury.Símbolo != '☿' {
		t.Errorf("Expected ☿, got %c", mercury.Símbolo)
	}
	if mercury.DiametroEcuatorial != 4878 {
		t.Errorf("Expected 4878, got %d", mercury.DiametroEcuatorial)
	}
	if mercury.PeríodoOrbital != 0.24 {
		t.Errorf("Expected 0.24, got %f", mercury.PeríodoOrbital)
	}
	if mercury.Inclinación != 7.0 {
		t.Errorf("Expected 7.0, got %f", mercury.Inclinación)
	}
}

// TestStructPartialInitialization tests partial initialization
func TestStructPartialInitialization(t *testing.T) {
	p := Planeta{DiametroEcuatorial: 4878}
	
	if p.DiametroEcuatorial != 4878 {
		t.Error("Specified field should have given value")
	}
	if p.Símbolo != 0 {
		t.Error("Unspecified field should have default value")
	}
	if p.PeríodoOrbital != 0.0 {
		t.Error("Unspecified field should have default value")
	}
}

// TestStructComparison tests struct comparison
func TestStructComparison(t *testing.T) {
	p1 := Planeta{Símbolo: '☿', DiametroEcuatorial: 4878}
	p2 := Planeta{Símbolo: '☿', DiametroEcuatorial: 4878}
	p3 := Planeta{Símbolo: '♀', DiametroEcuatorial: 12104}
	
	if p1 != p2 {
		t.Error("Structs with same values should be equal")
	}
	if p1 == p3 {
		t.Error("Structs with different values should not be equal")
	}
}

// TestStructFieldAccess tests accessing struct fields
func TestStructFieldAccess(t *testing.T) {
	mercury := Planeta{
		Símbolo:            '☿',
		DiametroEcuatorial: 4878,
		PeríodoOrbital:     0.24,
		Inclinación:        7.0,
	}
	
	// Read access
	diameter := mercury.DiametroEcuatorial
	if diameter != 4878 {
		t.Error("Field access failed")
	}
	
	// Write access
	mercury.Inclinación = 7.5
	if mercury.Inclinación != 7.5 {
		t.Error("Field modification failed")
	}
}

// TestStructWithSlice tests struct containing a slice
func TestStructWithSlice(t *testing.T) {
	earth := Planeta{Símbolo: '♁', DiametroEcuatorial: 12756}
	
	viaLáctea := Galaxia{
		Nombre:    "Via láctea",
		Velocidad: 220,
		Planetas:  []Planeta{earth},
	}
	
	if viaLáctea.Nombre != "Via láctea" {
		t.Error("String field failed")
	}
	if len(viaLáctea.Planetas) != 1 {
		t.Error("Slice field failed")
	}
	if viaLáctea.Planetas[0].Símbolo != '♁' {
		t.Error("Nested struct access failed")
	}
}

// TestStructPositionalInitialization tests initialization without field names
func TestStructPositionalInitialization(t *testing.T) {
	earth := Planeta{
		'♁',
		12756,
		1.0,
		0.0,
	}
	
	if earth.Símbolo != '♁' {
		t.Error("Positional initialization failed for first field")
	}
	if earth.DiametroEcuatorial != 12756 {
		t.Error("Positional initialization failed for second field")
	}
	if earth.PeríodoOrbital != 1.0 {
		t.Error("Positional initialization failed for third field")
	}
	if earth.Inclinación != 0.0 {
		t.Error("Positional initialization failed for fourth field")
	}
}

// TestStructEmbedding tests anonymous fields and struct embedding
func TestStructEmbedding(t *testing.T) {
	mercury := PlanetaExtended{}
	mercury.Símbolo = '☿'
	mercury.DiametroEcuatorial = 4878
	
	// Test direct access to embedded struct fields
	mercury.X = 128
	mercury.Y = -17
	mercury.Z = 3049
	mercury.T = time.Date(2020, 5, 19, 11, 0, 0, 0, time.UTC)
	
	if mercury.X != 128 {
		t.Error("Anonymous field access failed")
	}
	if mercury.Y != -17 {
		t.Error("Anonymous field access failed")
	}
	if mercury.Z != 3049 {
		t.Error("Anonymous field access failed")
	}
	
	// Test access through explicit path
	if mercury.Point.X != 128 {
		t.Error("Explicit path access failed")
	}
}

// Universidad represents a university
type Universidad struct {
	Url    string
	Nombre string
	Biblioteca
}

// Biblioteca represents a library
type Biblioteca struct {
	Url string
}

// TestStructFieldShadowing tests field shadowing in embedded structs
func TestStructFieldShadowing(t *testing.T) {
	uc3m := Universidad{}
	uc3m.Url = "https://www.uc3m.es"
	uc3m.Nombre = "Universidad Carlos III de Madrid"
	
	// The outer Url should be set
	if uc3m.Url != "https://www.uc3m.es" {
		t.Error("Outer field should be set")
	}
	
	// Set the inner Biblioteca.Url
	uc3m.Biblioteca.Url = "https://www.uc3m.es/biblioteca"
	
	if uc3m.Biblioteca.Url != "https://www.uc3m.es/biblioteca" {
		t.Error("Inner field should be accessible via explicit path")
	}
	
	// Outer Url should not be affected
	if uc3m.Url != "https://www.uc3m.es" {
		t.Error("Outer field should remain unchanged")
	}
}

// Link represents a network link
type Link struct {
	Server
	Client
}

// Server represents a server
type Server struct {
	name string
	addr string
	port int
}

// Client represents a client
type Client struct {
	name string
}

// GetPosition retrieves planet position (pass by value)
func GetPosition(planeta Planeta, t time.Time) TimedPosition {
	// Simplified implementation
	return TimedPosition{}
}

// SetPosition sets planet position (pass by reference)
func SetPosition(planeta *Planeta, tp TimedPosition) {
	// In a real implementation, this would update the planet's position
}

// TestStructPassByValue tests passing structs by value
func TestStructPassByValue(t *testing.T) {
	mercury := Planeta{
		Símbolo:            '☿',
		DiametroEcuatorial: 4878,
		PeríodoOrbital:     0.24,
		Inclinación:        7.0,
	}
	
	originalDiameter := mercury.DiametroEcuatorial
	_ = GetPosition(mercury, time.Now())
	
	// Struct should not be modified (pass by value)
	if mercury.DiametroEcuatorial != originalDiameter {
		t.Error("Struct passed by value should not be modified")
	}
}

// TestStructPassByReference tests passing structs by reference
func TestStructPassByReference(t *testing.T) {
	mercury := Planeta{
		Símbolo:            '☿',
		DiametroEcuatorial: 4878,
	}
	
	tp := TimedPosition{
		Point: Point{X: 100, Y: 200, Z: 300},
		T:     time.Now(),
	}
	
	SetPosition(&mercury, tp)
	
	// Function received a pointer and could modify the struct
	// (though our implementation doesn't modify it)
}

// LinkedList represents a linked list node
type LinkedList struct {
	valor int
	next  *LinkedList
}

// TestStructSelfReference tests struct with self-reference
func TestStructSelfReference(t *testing.T) {
	node1 := LinkedList{valor: 1}
	node2 := LinkedList{valor: 2}
	node3 := LinkedList{valor: 3}
	
	node1.next = &node2
	node2.next = &node3
	
	if node1.next.valor != 2 {
		t.Error("Linked list traversal failed")
	}
	if node1.next.next.valor != 3 {
		t.Error("Linked list traversal failed")
	}
	if node3.next != nil {
		t.Error("Last node should point to nil")
	}
}

// TestStructZeroValue tests zero values in structs
func TestStructZeroValue(t *testing.T) {
	var p Planeta
	
	if p.Símbolo != 0 {
		t.Error("Zero value for rune field should be 0")
	}
	if p.DiametroEcuatorial != 0 {
		t.Error("Zero value for int field should be 0")
	}
	if p.PeríodoOrbital != 0.0 {
		t.Error("Zero value for float64 field should be 0.0")
	}
	if p.Inclinación != 0.0 {
		t.Error("Zero value for float64 field should be 0.0")
	}
}

// TestStructNesting tests nested structs
func TestStructNesting(t *testing.T) {
	tp := TimedPosition{
		Point: Point{X: 1, Y: 2, Z: 3},
		T:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	
	if tp.Point.X != 1 || tp.Point.Y != 2 || tp.Point.Z != 3 {
		t.Error("Nested struct access failed")
	}
	if tp.T.Year() != 2020 {
		t.Error("Nested struct time field failed")
	}
}

// TestStructModification tests modifying struct fields
func TestStructModification(t *testing.T) {
	p := Planeta{DiametroEcuatorial: 1000}
	
	p.DiametroEcuatorial = 2000
	if p.DiametroEcuatorial != 2000 {
		t.Error("Struct field modification failed")
	}
	
	p.Símbolo = '♁'
	if p.Símbolo != '♁' {
		t.Error("Struct field modification failed")
	}
}

// TestStructInSlice tests slices of structs
func TestStructInSlice(t *testing.T) {
	planets := []Planeta{
		{Símbolo: '☿', DiametroEcuatorial: 4878},
		{Símbolo: '♀', DiametroEcuatorial: 12104},
		{Símbolo: '♁', DiametroEcuatorial: 12756},
	}
	
	if len(planets) != 3 {
		t.Errorf("Expected 3 planets, got %d", len(planets))
	}
	if planets[0].Símbolo != '☿' {
		t.Error("First planet should be Mercury")
	}
	if planets[2].DiametroEcuatorial != 12756 {
		t.Error("Third planet diameter should be 12756")
	}
}
