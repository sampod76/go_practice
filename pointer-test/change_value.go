package main

import "fmt"

func changeNumber(number *int) {
	// এখানে pointer দিয়ে original value change হবে
	*number = 100
}

func main() {
	result := 20
	changeNumber(&result)
	fmt.Println(result)
}
