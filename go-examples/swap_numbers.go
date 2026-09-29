package main

import "fmt"

func main() {
	first := 10
	second := 20

	// Go te temporary variable charao swap kora jay
	first, second = second, first
	fmt.Println(first, second)
}
