package main

import "fmt"

func calculate(a int, b int) (int, int) {
	// multiple value return practice
	return a + b, a - b
}

func main() {
	sum, difference := calculate(20, 8)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}
