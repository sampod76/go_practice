package main

import "fmt"

func main() {
	prices := map[string]int{
		"pen":  10,
		"book": 50,
	}

	// key exist kore kina check
	price, found := prices["book"]
	if found {
		fmt.Println("Price:", price)
	} else {
		fmt.Println("Item not found")
	}
}
