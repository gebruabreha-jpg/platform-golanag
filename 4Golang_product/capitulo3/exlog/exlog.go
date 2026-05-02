// Ejemplo del uso de servicios básicos de log
package main

import "log"

func main() {
	log.Println("Inicio del programa")
	log.Printf("Procesando %d elementos\n", 42)
	log.Fatalln("Error crítico")
}
