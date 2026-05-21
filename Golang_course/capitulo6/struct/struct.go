// -*- coding: utf-8 -*-
// struct.go
// -----------------------------------------------------------------------------
//
// Started on <mar 24-02-2026 06:31:36.803818448 (1771911096)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Ejemplo del uso de structs
package main

import (
	"fmt"

	"github.com/clinaresl/struct/dial"
)

type Universidad struct {
	Url    string
	Nombre string
	Biblioteca
}
type Biblioteca struct {
	Url string
}

func main() {

	// Ocultación de campos anónimos
	uc3m := Universidad{}
	uc3m.Url = "https://www.uc3m.es"
	uc3m.Nombre = "Universidad Carlos III de Madrid"
	fmt.Printf("%v\n", uc3m)
	uc3m.Biblioteca.Url = "https://www.uc3m.es/biblioteca"
	fmt.Printf("%v\n", uc3m)

	// Ordenación de elementos químicos
	items := []dial.Átomo{
		{"Be", 9.012},
		{"H", 1.008},
		{"Li", 6.967},
		{"He", 4.003}}
	if err := dial.Sort(items); err == nil {
		for _, v := range items {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	} else {
		fmt.Println(err)
	}
}

// Local Variables:
// mode:go
// fill-column:80
// End:
