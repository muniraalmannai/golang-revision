// Instructions: Write a program that displays the number of arguments passed to it.
// This number will be followed by a newline ('\n').
// If there is no argument, the program displays 0 followed by a newline.
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	c := len(args)
	a := 0
	if c == 0 {
		z01.PrintRune('c')
	} else {
		digits := []rune{}
		for c > 0 {
			a = c % 10
			digits = append([]rune{rune(a + '0')}, digits...)
			c /= 10
		}
		for i := 0; i < len(digits); i++ {
			z01.PrintRune(digits[i])
		}
	}
	z01.PrintRune('\n')
}
