package main

import "fmt"

func main() {
	numbers := []int{10, 25, 30, 45}
	find := 30
	found := false

	for _, number := range numbers {
		if number == find {
			found = true
			break
		}
	}

	// value ta slice e ache kina
	fmt.Println("Found:", found)
}
