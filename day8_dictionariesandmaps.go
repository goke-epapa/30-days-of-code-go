package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
@link https://www.hackerrank.com/challenges/30-dictionaries-and-maps/problem
*/

/*
Sample Input
3
sam 99912222
tom 11122222
harry 12299933
sam
edward
harry
*/

/*
Sample output
sam=99912222
Not found
harry=12299933
*/

func main() {
	print := fmt.Printf
	scanner := bufio.NewScanner(os.Stdin)
	n, err := strconv.Atoi(input(scanner))

	if err != nil {
		panic("invalid input")
	}

	phoneBook := make(map[string]string, n)

	// Create Phone Book
	for i := 0; i < n; i++ {
		entry := input(scanner)
		split := strings.Split(entry, " ")
		phoneBook[split[0]] = split[1]
	}

	// Find entries
	for {
		name := input(scanner)

		if name == "" {
			break
		}

		if phone, ok := phoneBook[name]; ok {
			print("%s=%s\n", name, phone)
		} else {
			print("Not found\n")
		}
	}
}

func input(scanner *bufio.Scanner) string {
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
