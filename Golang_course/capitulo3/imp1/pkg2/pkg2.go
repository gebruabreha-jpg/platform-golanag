package pkg2

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Subsistema 2: activado ...")
}

func Translate(cadena string) string {

	runes := []rune(cadena)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
