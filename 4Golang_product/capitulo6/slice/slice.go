// -*- coding: utf-8 -*-
// slice.go
// -----------------------------------------------------------------------------
//
// Started on <jue 19-02-2026 16:13:49.766529677 (1771514029)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Ejemplos del uso de slices
package main

import (
	"fmt"

	"github.com/clinaresl/slice/dial"
)

// -1 si slice1<slice2; +1 si slice1>slice2 y 0 si slice1==slice2
func lessThan(slice1, slice2 []byte) int {
	items := len(slice1)
	if len(slice2) < len(slice1) {
		items = len(slice2)
	}
	for idx := 0; idx < items; idx++ {
		if slice1[idx] < slice2[idx] {
			return -1
		}
		if slice1[idx] > slice2[idx] {
			return +1
		}
	}
	if items == len(slice1) && items == len(slice2) {
		return 0
	}
	if items == len(slice1) {
		return -1
	}
	return +1
}

// Actualización del primer valor de un slice
func updateValue(s []int, value int) {

	// Es peligroso asumir que existe la posición #0
	s[0] = value
}

// Intento (fallido) de añadir un elemento a un slice
func addValue(s []int, value int) {
	s = append(s, value)
}

// Inserción de un elemento al final de un slice
func AddValue(s *[]int, value int) {
	*s = append(*s, value)
}

func main() {

	// Por defecto, un slice es nil
	var slice []int
	fmt.Println(slice == nil)

	// Comparación de slices
	slice1 := []byte{'g', 'o'}     // slice "go"
	var array2 = [2]byte{103, 111} // array "go"
	fmt.Println(lessThan(slice1, array2[:]))

	// Conversión explícita a runas
	cadena := "sin(θ)=cos(π/2-θ)"
	runas := []rune(cadena)
	fmt.Println(runas)

	// Inicialización por defecto de slices
	s1 := make([]bool, 2)
	s1[1] = true
	fmt.Printf("%v\n", s1)

	s2 := make([]int, 3)
	s2[1] = 1
	fmt.Printf("%v\n", s2)

	s3 := make([]string, 4)
	s3[1] = "Hi!"
	fmt.Printf("%v\n", s3) // los elementos vacíos no se muestran
	fmt.Printf("'%v' '%v' '%v' '%v'\n", s3[0], s3[1], s3[2], s3[3])

	// Actualización del primer valor de un slice
	sl0 := []int{1, 2, 3}
	updateValue(sl0, 4)
	fmt.Printf("Contents %v\n", sl0)

	// El paso por valor no sirve para modificar un slice
	sl1 := []int{}
	addValue(sl1, 1)
	fmt.Printf("Contents %v\n", sl1)

	// Pero el paso por referencia sí lo hace
	sl4 := []int{}
	AddValue(&sl4, 1)
	fmt.Printf("Contents %v\n", sl4)

	sl5 := make([]int, 0, 1)
	AddValue(&sl5, 2)
	fmt.Printf("Contents %v\n", sl5)

	// Copia de slices
	s4 := make([]int, 3)
	s4[1] = 1
	fmt.Println(s4)

	s5 := make([]int, 3)
	copy(s5, s2)
	s5[1] = 2
	fmt.Println(s4)
	fmt.Println(s5)

	s6 := []int{1, 2, 3, 4, 5}
	s7 := []int{}
	copy(s6, s7)
	fmt.Println(s6)
	copy(s7, s6)
	fmt.Println(s7)
	s8 := make([]int, 0, 10) // la capacidad no importa
	copy(s8, s6)
	fmt.Println(s8)

	// Escritura del valor por defecto
	var z = []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("z: %v (len: %v, cap: %v)\n", z, len(z), cap(z))
	clear(z)
	fmt.Printf("z: %v (len: %v, cap: %v)\n", z, len(z), cap(z))

	// Ordenación con el algoritmo de Dial
	numeros := []int{3, 9, 2, 1, 8, 9, 13, 17, 0, 4}
	fmt.Println(numeros)
	if err := dial.Sort(numeros); err != nil {
		fmt.Println(" error durante la ordenación")
	} else {
		fmt.Println(numeros)
	}
}

// Local Variables:
// mode:go
// fill-column:80
// End:
