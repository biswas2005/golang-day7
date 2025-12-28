package practice

import (
	"fmt"
)

func Character() {

	str := "Golang"
	count := 0
	for _, ch := range str {

		//prints each character
		fmt.Printf("%c", ch)

		//ASCII-Unicode value of each character
		fmt.Print("-->", ch, "\n")
		count++
	}
	//Counts the number of character
	fmt.Println("Count", count)

	fmt.Println("Reverse is", reverse(str))
	fmt.Println("Is Empty", empty(str))
}

// Reverse a string
func reverse(str string) string {

	rune := []rune(str)

	for i, j := 0, len(rune)-1; i < j; i, j = i+1, j-1 {
		rune[i], rune[j] = rune[j], rune[i]
	}
	return string(rune)

}

// Check if string empty
func empty(str string) bool {

	Empty := false
	if len(str) <= 0 {
		Empty = true

	}
	return Empty

}
