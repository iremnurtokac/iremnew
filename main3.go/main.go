package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "<>"
	for i := 0; i < 5; i++ {

		str1 = strings.Replace(str1, "<>", "<<>>", -1)
		str1 = strings.Replace(str1, "<>", "<> <>", -1)
		fmt.Println(str1)
	}

}
