package main

import "fmt"

func main() {
	marks := 72

	if marks >= 80 {
		fmt.Println("Grade A")
	} else if marks >= 70 {
		fmt.Println("Grade B")
	} else if marks >= 60 {
		fmt.Println("Grade C")
	} else if marks >= 40 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("Failed")
	}
}
