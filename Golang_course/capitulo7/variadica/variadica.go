// -*- coding: utf-8 -*-
// variadica.go
// -----------------------------------------------------------------------------
//
// Started on <mar 24-02-2026 06:56:51.241430640 (1771912611)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Description
package main

import (
	"bytes"
	"fmt"
	"os"
)

type color struct {
	R, G, B byte
}

type colorString struct {
	str string
	color
}

func filterEven(nums ...int) (result []int) {

	result = []int{}

	for _, n := range nums {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}

func prefix(c color) string {

	var result bytes.Buffer
	_, err := fmt.Fprintf(&result, "\u001b[38;2;%d;%d;%dm",
		c.R, c.G, c.B)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf error: %v\n", err)
	}
	return result.String()
}

func suffix() string {
	return "\u001b[0m"
}

func cprint(cstr ...colorString) string {

	output := ""
	for _, v := range cstr {
		var result bytes.Buffer
		_, err := fmt.Fprintf(&result, "%s%s%s",
			prefix(v.color), v.str, suffix())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fprint error: %v\n", err)
		}
		output += result.String()
	}
	return output
}

func main() {

	// Ejemplos del uso de funciones variadicas
	for _, number := range filterEven(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) {
		fmt.Printf("%v ", number)
	}
	fmt.Println()

	cadena := cprint(colorString{"Rojo", color{255, 0, 0}},
		colorString{"Verde", color{0, 255, 0}},
		colorString{"Azul", color{0, 0, 255}},
		colorString{"Amarillo", color{255, 255, 0}})
	fmt.Println(cadena)

	spec := []colorString{
		colorString{"Rojo", color{126, 0, 0}},
		colorString{"Verde", color{0, 126, 0}},
		colorString{"Azul", color{0, 0, 126}},
		colorString{"Amarillo", color{126, 126, 0}},
	}
	fmt.Println(cprint(spec...))

}

// Local Variables:
// mode:go
// fill-column:80
// End:
