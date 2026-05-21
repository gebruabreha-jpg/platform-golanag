// -*- coding: utf-8 -*-
// func.go
// -----------------------------------------------------------------------------
//
// Started on <dom 15-02-2026 15:05:35.861299957 (1771164335)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Ejemplos de definición de funciones
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os/exec"
	"strings"
)

func div1(a, b float64) (result float64) {
	if b == 0 {
		return math.NaN()
	}
	result = a / b
	return
}

func div2(x, y float64) (float64, error) {
	if y == 0 {
		return math.NaN(), errors.New("Overflow")
	}
	return x / y, nil
}

// Devuelve los términos (i-1)- e i-ésimo de la serie
// de Fibonacci respectivamente, a partir de Fib(0)=x
// y Fib(1)=y, para i>1
func fibonacci(x, y, i uint64) (uint64, uint64, error) {
	n1, n0, n := y, x+y, uint64(2)
	for n < i {
		n1, n0, n = n0, n1+n0, n+1
		if n0 < n1 {
			return 0, 0, errors.New("Overflow")
		}
	}
	return n1, n0, nil
}

// Devuelve una estimación del valor de la razón áurea como el
// cociente entre dos términos consecutivos de la serie de
// Fibonacci y una indicación de si son válidos
func phi(beforelast, last uint64, err error) (float64, error) {
	if err != nil {
		return math.NaN(), err
	}
	return float64(last) / float64(beforelast), nil
}

// invoca el comando sort sobre un conjunto de líneas dadas como
// un string con '\n'. Devuelve otra cadena con las líneas ordenadas
// y una indicación de error si lo hubiera
func ordenar(lineas string) (string, error) {
	cmd := exec.Command("sort")
	cmd.Stdin = strings.NewReader(lineas)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Error durante la ejecución de 'sort': %v", err)
	}
	return out.String(), nil
}

// intercambio (fallido) de los contenidos de dos variables enteras
func swap(a, b int) {
	a, b = b, a
}

// Devuelve los valores de F(0) y F(1) elegidos por el usuario, y el
// índice que se desea calcular de la serie de Fibonacci
func parseArgs(n0, n1, i *int) {

	// lectura de los valores numéricos
	flag.IntVar(n0, "n0", 0, "F(0)")
	flag.IntVar(n1, "n1", 1, "F(1)")
	flag.IntVar(i, "i", 2, "índice que se desea calcular, F (i)")

	flag.Parse()
}

// el procesamiento de los argumentos de usuario se puede hacer en init
func init() {
	var n0, n1, i int
	parseArgs(&n0, &n1, &i)
	if _, item, err := fibonacci(uint64(n0), uint64(n1), uint64(i)); err != nil {
		fmt.Printf(" Error durante la computación de Fibonacci: %v\n", err)
	} else {
		fmt.Printf("F(%v) = %v\n", i, item)
	}
}

func main() {

	// Prueba de la división entre dos números
	fmt.Println(div1(13, 4))
	fmt.Println(div1(13, 0))

	// demostración de shadowing
	const x = "interface{} says nothing"
	for x := range 3 {
		x := x * 2 // oculta el contador del bucle
		fmt.Printf("x: %v\n", x)
	}
	fmt.Printf("x: %v\n", x)

	// demostración del uso de funciones multi-valuadas
	fmt.Printf("10/0 = ")
	if res, err := div2(10, 0); err == nil {
		fmt.Printf("%v\n", res)
	} else {
		fmt.Printf("%v\n", err)
	}

	// demostración de la captura de valores de funciones multi-valuadas
	if razónÁurea, err := phi(fibonacci(0, 1, 50)); err == nil {
		fmt.Printf("Razón áurea: %v\n", razónÁurea)
	} else {
		fmt.Println(err)
	}

	var result string
	var err error
	if result, err = ordenar("Wilhelm Leibniz\nAlan Turing\nGottlob Frege"); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", result)

	// demostración del paso por valor (que no cambia los valores)
	op1, op2 := 2, 3
	fmt.Printf("op1: %v -- op2: %v\n", op1, op2)
	swap(op1, op2)
	fmt.Printf("op1: %v -- op2: %v\n", op1, op2)
}

// Local Variables:
// mode:go
// fill-column:80
// End:
