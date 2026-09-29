package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	numbers[0] = 5
	numbers[1] = 10
	numbers[2] = 15

	fmt.Println(numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}
