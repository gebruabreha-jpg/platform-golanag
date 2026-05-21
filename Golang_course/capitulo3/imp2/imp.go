// Ejercicio de inicialización de paquetes
package main

import (
	"fmt"

	"github.com/clinaresl/imp/payload"
	"github.com/clinaresl/imp/pkg2"
	_ "github.com/clinaresl/imp/pkg3"
)

func init() {
	fmt.Println("Main: Ready to launch ...")
}

func main() {
	payload.Ready()
	saludo := "Main: Liftoff!"
	fmt.Println(saludo)
	fmt.Println(pkg2.Translate("Arecibo message sent"))
}
