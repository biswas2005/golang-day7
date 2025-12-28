package learningstrings

import (
	"fmt"
	"strings"
)

func Trimm() {
	str := "      @@ Hello,World!!          "

	//simple trim
	fmt.Println(strings.Trim(str, "@ !"))

	//TrimLeft removes string from starting of string
	fmt.Println(strings.TrimLeft(str, "@ H"))

	//TrimRight removes string from ending of string
	fmt.Println(strings.TrimRight(str, "! l d"))

	//TrimSpace removes space at the start and end of string
	//spaces in between stays as it may be intentional
	fmt.Println(strings.TrimSpace(str))

	//TrimPrefix
	fmt.Println(strings.TrimPrefix(str, "      @"))

	//TrimSuffix
	fmt.Println(strings.TrimSuffix(str, "!!          "))
}
