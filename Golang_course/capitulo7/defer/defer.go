// -*- coding: utf-8 -*-
// defer.go
// -----------------------------------------------------------------------------
//
// Started on <mar 24-02-2026 07:12:22.108617269 (1771913542)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Ejemplos de uso de funciones diferidas
package main

import (
	"log"
)

func caso1() {
	i := 1
	defer log.Printf("i: %v\n", i)
	i++
}

func caso2() {
	i := 1
	defer log.Printf("i: %v\n", i)
	i++
	defer log.Printf("i: %v\n", i)
}

func main() {

	caso1()
	caso2()
}

// Local Variables:
// mode:go
// fill-column:80
// End:
