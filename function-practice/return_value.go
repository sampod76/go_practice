package main

import "fmt"

func multiply(a int, b int) int {
	// result return korbe
	return a * b
}

func main() {
	result := multiply(5, 4)
	fmt.Println(result)
}
