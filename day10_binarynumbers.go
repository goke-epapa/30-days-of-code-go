package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int64(nTemp)
	str := strconv.FormatInt(n, 2)

	var total, count int32
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "1" {
			count++
			if count > total {	
				total = count
			}
		} else {
			count = 0
		}
	}
	fmt.Println(total)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
