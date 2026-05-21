// Este modulo traduce cadenas de texto al complicado
// lenguaje Klingon
package pkg2

import (
	"fmt"
	"math/rand"

	_ "github.com/clinaresl/imp/payload"
)

func init() {
	fmt.Println("Módulo2: activado ...")
}

// traduce cualquier cadena a Klingon, lo creas o no
func Translate(cadena string) string {

	runes := []rune(cadena)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
