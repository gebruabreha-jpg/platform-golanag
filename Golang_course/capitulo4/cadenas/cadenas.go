// -*- coding: utf-8 -*-
// cadenas.go
// -----------------------------------------------------------------------------
//
// Started on <vie 06-02-2026 12:10:45.682258158 (1770376245)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Description
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	caracter := "Go"[0] // G: byte 0x47
	s := "G' Day"
	rev := s[3:] + s[:3] // "DayG' "
	fmt.Println(caracter)
	fmt.Println(rev)

	str := "sin(θ)=cos(π/2-θ)"
	fmt.Printf("%v\n", len(str))
	fmt.Printf("%v\n", utf8.RuneCountInString(str))
	fmt.Printf("% x\n", str)

	fmt.Printf("%x (%c)", str[2], str[2]) // 6e (n)
	substr := str[3:5]
	fmt.Println(substr) // (·\ucr·

	pos := 0
	fórmula := []byte("sin(θ)=cos(π/2-θ)")

	for len(fórmula) > 0 {
		nextrun, tamaño := utf8.DecodeRune(fórmula)
		fmt.Printf("%c %2v %v\n", nextrun, pos, tamaño)

		pos += tamaño
		fórmula = fórmula[tamaño:]
	}
	fmt.Println()

	for pos, runa := range "sin(θ)=cos(π/2-θ)" {
		fmt.Printf("%c %2v\n", runa, pos)
	}

	reg := `^\s*([A-Za-z0-9]+)\s*:\s*%([A-Za-z]+)\s*`
	fmt.Println(reg)

	fmt.Printf("\x47\x6f\n") // Go
	fmt.Printf("\u2655\n")   // dama blanca de ajedrez
}

// Local Variables:
// mode:go
// fill-column:80
// End:
