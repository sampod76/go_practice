package main

import "fmt"

func main() {
	age := 20
	hasCard := true

	if age >= 18 {
		// first condition true, ekhon card check korbo
		if hasCard {
			fmt.Println("Entry allowed")
		}
	}
}
