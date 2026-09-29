package main

import "fmt"

func main() {
	word := "golang"
	count := 0

	for _, letter := range word {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}

	fmt.Println("Vowels:", count)
}
