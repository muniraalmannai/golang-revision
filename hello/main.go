// Instructions: Write a program that displays "Hello World!" followed by a newline ('\n').
package main

import "github.com/01-edu/z01"

func main() {
	s := "Hello World!"

	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}
