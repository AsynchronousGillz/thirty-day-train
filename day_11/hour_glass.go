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
	var sum int32 = -69
	var temp int32 = 0
	for i := 1; i < len(arr) - 1; i++ {
		for j := 1; j < len(arr[i]) - 1; j++ {
			temp = arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] + arr[i][j] + arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]
			//fmt.Printf("%d %d %d\n", arr[i-1][j-1], arr[i-1][j], arr[i-1][j+1])
			//fmt.Printf("  %d  \n", arr[i][j])
			//fmt.Printf("%d %d %d\n", arr[i+1][j-1], arr[i+1][j], arr[i+1][j+1])
			if temp > sum {
				sum = temp
				temp = 0
			}
		} 
	}
	fmt.Printf("%d", sum)
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

