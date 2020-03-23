package main

/*
Given the meal price (base cost of a meal), tip percent (the percentage of the meal price being added as tip), and tax percent (the percentage of the meal price being added as tax) for a meal, find and print the meal's total cost.
*/

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Only this method was written, the rest of the code is from Hackerrank
func solve(mealCost float64, tipPercent int32, taxPercent int32) {
	tip := float64(tipPercent) * mealCost / 100
	tax := float64(taxPercent) * mealCost / 100

	mealTotalCost := math.Round(int64(mealCost + tip + tax))
	fmt.Println(mealTotalCost)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mealCost, err := strconv.ParseFloat(readLine(reader), 64)
	checkError(err)

	tipPercentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tipPercent := int32(tipPercentTemp)

	taxPercentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	taxPercent := int32(taxPercentTemp)

	solve(mealCost, tipPercent, taxPercent)
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
