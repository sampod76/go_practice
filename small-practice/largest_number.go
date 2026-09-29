package main

import "fmt"

func main() {
	a := 25
	b := 42
	c := 18
	largest := a

	// b and c er sathe compare korchi
	if b > largest {
		largest = b
	}
	if c > largest {
		largest = c
	}

	fmt.Println("Largest:", largest)
}
