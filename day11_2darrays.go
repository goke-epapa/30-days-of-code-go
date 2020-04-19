package main

/**
Calculate the hourglass sum for every hourglass in A, then print the maximum hourglass sum.
Input
1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0

2 4 4
  2
1 2 4
Output: 19
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



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }
    
    var maxSum int64 = math.MinInt64
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            sum := sum2dSlice(getHourGlass(i, j, arr))

            if sum > maxSum {
                maxSum = sum
            }
        }
    }

    fmt.Println(strconv.FormatInt(maxSum, 10))
}

func sum2dSlice(arr [][]int32) int64 {
    var sum int64 = 0
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr[i]); j++ {
            sum += int64(arr[i][j])
        }
    }

    return sum
}

func getHourGlass(rowIndex int, colIndex int, arr [][]int32) [][]int32 {
    fallbackHourGlass := [][]int32{{0, 0, 0},{0, 0, 0},{0, 0, 0}}
    lastColumnIndex := colIndex + 2
    lastRowIndex := rowIndex + 2

    if len(arr) == 0 || lastColumnIndex >= len(arr) || lastRowIndex >= len(arr[0]) {
        return fallbackHourGlass     
    }

    hourGlass := make([][]int32, 3)
    hourGlass[0] = arr[rowIndex][colIndex:lastColumnIndex + 1]
    hourGlass[1] = []int32{0, arr[rowIndex + 1][colIndex + 1], 0}
    hourGlass[2] = arr[lastRowIndex][colIndex:lastColumnIndex + 1]

    return hourGlass
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
