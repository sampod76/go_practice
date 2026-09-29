package main

import "fmt"

func main() {
	numbers := []int{5, 10, 15, 20}
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	fmt.Println("Sum:", sum)
}
