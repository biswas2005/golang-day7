package learningstrings

import (
	"fmt"
	"strings"
)

func Compare() {

	s1 := "Hello"
	s2 := "Abhi"
	s3 := "Hello"

	//comparision operators
	fmt.Println("s1==s2", s1 == s2)
	fmt.Println("s1==s3", s1 == s3)
	fmt.Println("s1!=s2", s1 != s2)

	//lexicographical comparision
	fmt.Println("s1<s3", s3 < s1)
	fmt.Println("s1>s2", s1 > s2)
	fmt.Println("s1<=s2", s1 <= s2)
	fmt.Println("s1>=s2", s1 >= s2)

	//strings.Compare
	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s3))
	fmt.Println(strings.Compare(s2, s3))

}
