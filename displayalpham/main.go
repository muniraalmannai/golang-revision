// Instructions: Write a program that displays the alphabet, with even letters in uppercase,
// and odd letters in lowercase, followed by a newline ('\n')
package main

import "github.com/01-edu/z01"

func main() {
	for i := 97; i < 123; i++ {
		if i%2 != 0 {
			z01.PrintRune(rune(i))
		} else {
			z01.PrintRune(rune(i - 32))
		}
	}
	z01.PrintRune('\n')
}
