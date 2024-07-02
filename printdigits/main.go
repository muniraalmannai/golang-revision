// Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.
// A line is a sequence of characters preceding the end of line character ('\n').
package main

import "github.com/01-edu/z01"

func main() {
	s := "0123456789"

	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}
