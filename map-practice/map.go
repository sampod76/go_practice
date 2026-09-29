package main

import "fmt"

func main() {
	studentMarks := map[string]int{
		"Rahim": 80,
		"Karim": 75,
	}

	// map theke value get kortesi
	fmt.Println(studentMarks["Rahim"])
}
