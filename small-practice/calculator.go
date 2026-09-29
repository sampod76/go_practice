package main

import "fmt"

func main() {
	firstNumber := 12.0
	secondNumber := 4.0
	operator := "*"

	switch operator {
	case "+":
		fmt.Println(firstNumber + secondNumber)
	case "-":
		fmt.Println(firstNumber - secondNumber)
	case "*":
		fmt.Println(firstNumber * secondNumber)
	case "/":
		fmt.Println(firstNumber / secondNumber)
	default:
		fmt.Println("Invalid operator")
	}
}
