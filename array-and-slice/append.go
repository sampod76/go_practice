package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}

	// slice er last e new value add
	numbers = append(numbers, 40)
	fmt.Println(numbers)
}
