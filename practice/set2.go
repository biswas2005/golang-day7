package practice

import (
	"fmt"
	"strings"
)

func Set2() {
	str := "madam"
	vowels := "aeiouAEIOU"
	vowelCount := 0
	consonantCount := 0

	for _, ch := range str {
		if strings.ContainsRune(vowels, ch) {
			vowelCount++
		} else {
			consonantCount++
		}
	}

	//count of vowels
	fmt.Println("Count of Vowels:", vowelCount)

	//count of consonents
	fmt.Println("COunt of Consonants:", consonantCount)

	//calls the revers func()
	if revers(str) == str {
		fmt.Println(str, "Is a palindrome.")

	} else {
		fmt.Println(str, "Is not a palindrome.")
	}

	//calls the words func()
	fmt.Println("There are\"", words("Hello how are you"), "\"words")

	//calls the replace func()
	fmt.Println(replace("Hello how are you"))

}

// revers() to check if string is palindrome
func revers(s string) string {

	rune := []rune(s)
	for i, j := 0, len(rune)-1; i < j; i, j = i+1, j-1 {
		rune[i], rune[j] = rune[j], rune[i]
	}
	return string(rune)
}

// words() counts the words in sentence
func words(s string) int {
	Words := strings.Split(s, " ")

	return len(Words)

}

// replace() to replace " " with "-"
func replace(s string) string {

	Result := strings.ReplaceAll(s, " ", "-")
	return Result
}
