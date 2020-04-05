package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
@link https://www.hackerrank.com/challenges/30-review-loop/problem
*/

func main() {
	print := fmt.Printf
	reader := bufio.NewScanner(os.Stdin)

	n, err := strconv.Atoi(input(reader))

	if err != nil {
		panic("Invalid input entered")
	}

	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = input(reader)
	}

	for _, str := range s {
		var sEven, sOdd string
		for j := 0; j < len(str); j++ {
			if j%2 == 0 {
				sEven += string(str[j])
			} else {
				sOdd += string(str[j])
			}
		}
		print("%s %s\n", sEven, sOdd)
	}
}

func input(reader *bufio.Scanner) string {
	if reader.Scan() {
		return reader.Text()
	}
	return ""
}
