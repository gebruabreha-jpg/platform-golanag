// -*- coding: utf-8 -*-
// if.go
// -----------------------------------------------------------------------------
//
// Started on <sáb 14-02-2026 19:28:36.936849245 (1771093716)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Ejemplos del uso de if
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	n := 1
	if n == 0 { // en vez de "if !n"
		log.Fatal("Debe proporcionar un valor numérico")
	}

	año := 2020
	añoBisiesto := año%4 == 0 && (año%100 != 0 ||
		(año%100 == 0 && año%400 == 0))
	if añoBisiesto {
		fmt.Printf(" El año %v es bisiesto\n", año)
	} else {
		fmt.Printf(" El año %v no es bisiesto\n", año)
	}

	contents := []byte("Curso de Go - Módulo I")
	if err := os.WriteFile("curso.txt", contents, 0644); err != nil {
		log.Fatal("Error de escritura")
	}
}

// Local Variables:
// mode:go
// fill-column:80
// End:
