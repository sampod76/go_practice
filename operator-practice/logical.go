package main

import "fmt"

func main() {
	age := 20
	hasID := true

	fmt.Println("Can enter:", age >= 18 && hasID)
	fmt.Println("Need checking:", age < 18 || !hasID)
}
