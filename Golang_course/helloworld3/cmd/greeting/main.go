package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Guest", "Your name")
	lang := flag.String("lang", "en", "Language (en, am, es)")
	flag.Parse()

	greetings := map[string]string{
		"en": "Hello",
		"am": "ሰላም",
		"es": "Hola",
	}

	greeting := greetings[*lang]
	if greeting == "" {
		greeting = greetings["en"]
	}

	fmt.Printf("%s, %s! 👋\n", greeting, *name)
}
