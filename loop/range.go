package main

import "fmt"

func main() {
	fruits := []string{"apple", "mango", "orange"}

	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}
}
