// -*- coding: utf-8 -*-
// map.go
// -----------------------------------------------------------------------------
//
// Started on <lun 23-02-2026 21:31:59.358903442 (1771878719)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Ejemplos del uso de mapas
package main

import (
	"fmt"

	"github.com/clinaresl/mod/multimap"
)

func main() {

	// Creación de un mapa vacío
	var mapa1 map[int]complex128
	fmt.Println(mapa1)

	var s []int
	var mapa2 map[int]complex128
	fmt.Printf("%v %v\n", s == nil, mapa2 == nil)

	// Inicialización de mapas
	demografía1 := map[string]int{
		"España":   46.94e+6,
		"Portugal": 10.28e+6,
		"Francia":  66.99e+6,
	}
	for k, v := range demografía1 {
		fmt.Printf("La población de %v es %v habitantes\n", k, v)
	}
	fmt.Println()

	var demografía2 map[string]int = make(map[string]int)
	demografía2["España"] = 46.94e+6
	demografía2["Portugal"] = 10.28e+6
	demografía2["Francia"] = 66.99e+6
	for k, v := range demografía2 {
		fmt.Printf("La población de %v es %v habitantes\n", k, v)
	}

	// Acceso de posiciones no existentes
	fmt.Printf("La población de Italia es %v habitantes\n",
		demografía2["Italia"])

	if value, ok := demografía2["Italia"]; !ok {
		fmt.Println("Población desconocida")
	} else {
		fmt.Printf("La población de Italia es %v\n", value)
	}

	// Eliminación de elementos de un mapa
	delete(demografía2, "España")
	delete(demografía2, "Italia")
	for k, v := range demografía2 {
		fmt.Printf("La población de %v es %v habitantes\n", k, v)
	}

	// Eliminación de las entradas de un mapa
	a := map[int]bool{1: true, 2: true, 3: true, 4: false, 5: true, 6: false}
	fmt.Printf("a: % v / len: %d\n", a, len(a))
	clear(a)
	fmt.Printf("a: % v / len: %d\n", a, len(a))
	fmt.Printf("[3]: %v\n", a[3])

	// Ejemplo del uso de multimaps de cadenas a enteros
	grandSlams := make(map[string][]int)
	multimap.Set("Novak Djokovic", 10, grandSlams)
	multimap.Set("Novak Djokovic", 3, grandSlams)
	multimap.Set("Novak Djokovic", 7, grandSlams)
	multimap.Set("Novak Djokovic", 4, grandSlams)

	multimap.Set("Rafael Nadal", 2, grandSlams)
	multimap.Set("Rafael Nadal", 14, grandSlams)
	multimap.Set("Rafael Nadal", 2, grandSlams)
	multimap.Set("Rafael Nadal", 4, grandSlams)

	if values, valid := multimap.Get("Rafael Nadal", grandSlams); !valid {
		fmt.Println("Rafael Nadal no ha ganado nunca?")
	} else {
		for _, nbslams := range values {
			fmt.Printf("Grand Slams ganados por Rafael Nadal: %v\n", nbslams)
		}
	}

	multimap.Remove("Novak Djokovic", 10, grandSlams)
	multimap.Remove("Rafael Nadal", 14, grandSlams)
	if values, valid := multimap.Get("Rafael Nadal", grandSlams); !valid {
		fmt.Println("Rafael Nadal no ha ganado nunca?")
	} else {
		for _, nbslams := range values {
			fmt.Printf("Grand Slams ganados por Rafael Nadal: %v\n", nbslams)
		}
	}
}

// Local Variables:
// mode:go
// fill-column:80
// End:
