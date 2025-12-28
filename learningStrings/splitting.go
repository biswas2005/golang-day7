package learningstrings

import (
	"fmt"
	"strings"
)

func Splitt() {

	str := "Hello,welcome,how,are,you"

	//Split divides on every key found
	fmt.Println(strings.Split(str, "e"))

	//splitAfter wont remove the key and divide
	fmt.Println(strings.SplitAfter(str, ","))

	//SplitN divides into N slices
	fmt.Println(strings.SplitN(str, ",", 3))

	//SplitAfterN divides into N slices and won't remove the key
	fmt.Println(strings.SplitAfterN(str, ",", 3))
}
