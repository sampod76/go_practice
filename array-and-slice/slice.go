package main

import "fmt"

func main() {
	fruits := []string{"apple", "mango", "banana"}

	fruits[1] = "orange"
	fmt.Println(fruits)
}
