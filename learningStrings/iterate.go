package learningstrings

import (
	"fmt"
	"unicode/utf8"
)

func Iteration() {

	//Iterate over a string
	//if %c not used to store each char
	//it prints the UTF-8 encoded points
	for index, str := range "Abhi Biswas" {
		fmt.Printf("The index for %c is %d:\n", str, index)
	}

	//access individual bytes of character
	//%x to print byte of character
	str := "Hello, Abhi"
	for i := 0; i < len(str); i++ {
		fmt.Printf("For character %c bytes are %x\n", str[i], str[i])
	}

	//To print the length we can
	//use len()
	//runecount returns number of rune present
	mystr := "Welcome, to GO!"
	fmt.Println("String is", mystr)
	fmt.Println("Method 1:", len(mystr))
	fmt.Println("Method 2:", utf8.RuneCountInString(mystr))

}
