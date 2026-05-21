// -*- coding: utf-8 -*-
// array.go
// -----------------------------------------------------------------------------
//
// Started on <dom 15-02-2026 16:49:32.875876326 (1771170572)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Ejemplos del uso de arrays
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"unicode"
)

func capitalizeByValue(cadena [2]byte) {
	cadena[0] = byte(unicode.ToUpper(rune(cadena[0])))
}

func capitalizeByReference(cadena *[2]byte) {
	(*cadena)[0] = byte(unicode.ToUpper(rune(cadena[0])))
}

func main() {

	array1 := [2]float64{}
	array2 := [2]float64{}
	array3 := [3]float64{}
	fmt.Println(reflect.TypeOf(array1) == reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(array1) == reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(array2) == reflect.TypeOf(array3))
	fmt.Println(array1 == array2)

	// Inicialización con expresiones literales
	var a = [2]int{2, 4}
	b := [3]int{6, 8, 10}
	fmt.Println(a)
	fmt.Println(b)

	// inicialización explícita con menos items
	c := [3]byte{'g', 'o'}
	d := [2]complex128{}
	fmt.Println(c)
	fmt.Println(d)

	// inicialización de un array con un literal
	e := [...]byte{'g', 'o'}
	fmt.Println(e)

	// inicialización selectiva de arrays
	cuadrados := [...]int{1: 1, 2: 4, 3: 9, 4: 16}
	fmt.Println(cuadrados)

	const (
		España = iota
		Portugal
		Francia
	)
	capitales := [...]string{España: "Madrid",
		Portugal: "Lisboa",
		Francia:  "París"}
	for k, v := range capitales {
		fmt.Printf("La capital de %v es %v\n", k, v)
	}

	// Generación de 10 números enteros aleatorios en el
	// intervalo [0, 10)
	random := [10]int{}

	// En Go 1.20+, rand se inicializa automáticamente
	for k := range 10 {
		random[k] = rand.Intn(10)
	}
	fmt.Println(random)

	// pruebas del paso por valor y por referencia
	array := [...]byte{'g', 'o'}

	capitalizeByValue(array)
	fmt.Println(string(array[:]))

	capitalizeByReference(&array)
	fmt.Println(string(array[:]))

}

// Local Variables:
// mode:go
// fill-column:80
// End:
