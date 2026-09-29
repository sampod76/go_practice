package main

import "fmt"

func main() {
	age := 21
	agePointer := &age

	fmt.Println("Value:", age)
	fmt.Println("Address:", agePointer)
	fmt.Println("Value from pointer:", *agePointer)
}
