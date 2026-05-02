// -*- coding: utf-8 -*-
// for.go
// -----------------------------------------------------------------------------
//
// Started on <sáb 14-02-2026 20:18:09.490631775 (1771096689)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Ejemplos del uso de for
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// Conversión de números binarios
	valor := "00110010"
	resultado := 0
	for idx := 0; idx < 8; idx++ {
		resultado <<= 1
		if valor[idx] == '1' {
			resultado += 1
		}
	}
	fmt.Printf("%s: %v\n", valor, resultado)

	final := make([]byte, 12)
	s := "Roma es Amor"
	// inversion de una cadena 's'
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		final[i], final[j] = s[j], s[i]
	}

	// calcular la suma de todos los números naturales hasta el primero que sea
	// múltiplo de 42, y que no sean múltiplos ni de 3 ni de 5
	suma := 0
	for i := 1; ; i++ { // sin condición de terminación!
		if i%42 == 0 {
			break
		}
		if i%3 == 0 || i%5 == 0 {
			continue
		}
		suma += i
	}
	fmt.Println("Suma: ", suma)

	// programación del comando UNIX cat
	filename := "curso.txt" // creado en la sección 2 (if)
	stream, err := os.Open(filename)
	if err != nil {
		log.Fatalf(" Error abriendo el fichero: %v", err)
	}
	reader := bufio.NewScanner(stream)
	for reader.Scan() {
		fmt.Printf("%v\n", reader.Text())
	}

	// Versión mejorada del traductor binario
	resultado = 0
	for idx := range 8 {
		resultado <<= 1
		if valor[idx] == '1' {
			resultado += 1
		}
	}
	fmt.Printf("%s: %v\n", valor, resultado)

	// emulación de string.ToUpper
	proverb := "Make the zero value useful!"
	for _, v := range proverb {
		if v >= 97 && v <= 122 { // es una letra minúscula
			v -= 32 // conversión a mayúscula
		}
		fmt.Printf("%c", v) // muestra el caracter i-ésimo
	}
	fmt.Println()
}

// Local Variables:
// mode:go
// fill-column:80
// End:
