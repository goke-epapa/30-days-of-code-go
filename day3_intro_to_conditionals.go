package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

Given an integer, n, perform the following conditional actions:

If n is odd, print Weird
If n is even and in the inclusive range of 2 to 5, print Not Weird
If n is even and in the inclusive range of 6 to 20, print Weird
If n is even and greater than 20, print Not Weird
Complete the stub code provided in your editor to print whether or not  is weird.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	i, err := strconv.Atoi(strings.TrimRight(input, "\n"))

	if err != nil {
		panic(err)
	}

	var out string
	if i%2 != 0 || i >= 6 && i <= 20 {
		out = "Weird"
	} else {
		out = "Not Weird"
	}

	fmt.Println(out)
}
