// Instructions: Write a function that counts the runes of a string and that returns that count.
package main

import "fmt"

func StrLen(s string) int {
	a := len(s)
	return a
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
