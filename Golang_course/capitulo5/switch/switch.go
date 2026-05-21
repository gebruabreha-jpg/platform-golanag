// -*- coding: utf-8 -*-
// switch.go
// -----------------------------------------------------------------------------
//
// Started on <sáb 14-02-2026 19:51:44.638781516 (1771095104)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Ejemplos de uso de switch
package main

import "fmt"

func main() {

	val := 18
	switch {
	case val%2 == 0:
		fmt.Println("es múltiplo de 2")
	case val%3 == 0:
		fmt.Println("es múltiplo de 3")
	case val%6 == 0:
		fmt.Println("es múltiplo de 6")
	}

	jaque := false
	movimientos := 42
	switch {
	case jaque && movimientos == 0:
		fmt.Println("Jaque mate\n")
	case !jaque && movimientos == 0:
		fmt.Printf("Ahogado\n")
	case jaque && movimientos > 0:
		fmt.Printf("Jaque\n")
	default:
		fmt.Printf("Su turno\n")
	}

	var dias int
	añoBisiesto := false
	mes := "Agosto"
	switch mes {
	case "Enero", "Marzo", "Mayo", "Julio", "Agosto", "Octubre", "Diciembre":
		dias = 31
	case "Abril", "Junio", "Septiembre", "Noviembre":
		dias = 30
	case "Febrero":
		dias = 28
		if añoBisiesto {
			dias = 29
		}
	}
	fmt.Printf("%v tiene %v días\n", mes, dias)

	pieza := "Dama"
	switch pieza {
	case "Peón":
		fmt.Println("Calculando movimientos de peón")
	case "Caballo":
		fmt.Println("Calculando movimientos en L")
	case "Alfil", "Dama":
		fmt.Println("Calculando movimientos diagonales")
		if pieza == "Alfil" {
			break // fallthrough debe estar en el primer nivel
		}
		fallthrough
	case "Torre": // La Dama ejecuta el caso anterior y este
		fmt.Println("Calculando movimientos horizontales")
	case "Rey":
		fmt.Println("Calculando movimientos de un paso")
	}

}

// Local Variables:
// mode:go
// fill-column:80
// End:
