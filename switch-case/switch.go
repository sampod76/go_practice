package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	default:
		fmt.Println("Another day")
	}
}
