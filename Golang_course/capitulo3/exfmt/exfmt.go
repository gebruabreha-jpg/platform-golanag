// Ejemplo del uso de servicios básicos de impresión
// de fmt
package main

import "fmt"

func main() {
	fmt.Println("Hola desde", "Go")
	fmt.Printf("Versión: %.2f\n", 1.25)
	fmt.Printf("Tipo: %T, Valor: %v\n", "Go", "Go")
}

// Hola desde Go
// Versión: 1.25
// Tipo: string, Valor: Go
