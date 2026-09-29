package main

import "fmt"

func main() {
	day := "Friday"

	switch day {
	case "Friday", "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Working day")
	}
}
