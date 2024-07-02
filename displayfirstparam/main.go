// Instructions: Write a program that displays its first argument, if there is one.
package main

import (
	"os"
	"github.com/01-edu/z01"
	)

func main(){
	args := os.Args
	if len(args) < 2 {
		return
	}
	for _, c := range args[1] {
		z01.PrintRune(rune(c))
	}
z01.PrintRune('\n')
}