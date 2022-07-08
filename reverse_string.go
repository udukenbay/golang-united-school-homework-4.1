package reverse_string

//package main

import (
	"fmt"
)

func main() {
	fmt.Println(ReverseString("Hello world!"))
}

func ReverseString(input string) (output string) {
	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}

	return output
}
