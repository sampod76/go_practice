package main

import "fmt"

func main() {
	temperature := 28

	if temperature > 30 {
		fmt.Println("Hot")
	} else if temperature >= 20 {
		fmt.Println("Normal")
	} else {
		fmt.Println("Cold")
	}
}
