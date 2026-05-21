package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()
	var greeting string

	switch {
	case hour < 12:
		greeting = "Good Morning"
	case hour < 18:
		greeting = "Good Afternoon"
	default:
		greeting = "Good Evening"
	}

	fmt.Printf("%s, World! 🌍\n", greeting)
	fmt.Printf("Current time: %s\n", time.Now().Format("15:04:05"))
}
