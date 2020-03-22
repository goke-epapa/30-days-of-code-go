package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
To complete this challenge, you must save a line of input from stdin to a variable, print Hello, World. on a single line, and finally print the value of your variable on a second line.
*/

func main() {
	fmt.Println("Hello, World.")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}