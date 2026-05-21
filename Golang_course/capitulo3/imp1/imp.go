// Ejercicio de inicialización de paquetes
package main

import (
	"fmt"

	"github.com/clinaresl/imp/pkg2"
)

func init() {
	fmt.Println("Main: Ready to launch ...")
}

func main() {
	saludo := "Main: Liftoff!"
	fmt.Println(saludo)
	fmt.Println(pkg2.Translate("Arecibo message sent"))
}
