package main

import "fmt"

type Student struct {
	name       string
	department string
	semester   int
}

func main() {
	student := Student{"Amin", "CSE", 4}
	fmt.Println("Name:", student.name)
	fmt.Println("Department:", student.department)
	fmt.Println("Semester:", student.semester)
}
