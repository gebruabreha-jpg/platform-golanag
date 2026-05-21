package main

import "fmt"

func main() {
	result := helperFunction(5) // Depends on helperFunction from helper.go
	fmt.Println("Result:", result)
}
