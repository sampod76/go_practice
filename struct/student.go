package main

import "fmt"

type Student struct {
	name  string
	age   int
	marks int
}

func main() {
	// struct diye student er data rakhlam
	student := Student{name: "Rima", age: 20, marks: 82}
	fmt.Println(student.name, student.age, student.marks)
}
