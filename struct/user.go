package main

import "fmt"

type User struct {
	username string
	email    string
	active   bool
}

func main() {
	user := User{
		username: "sajib",
		email:    "sajib@example.com",
		active:   true,
	}

	fmt.Println(user)
}
