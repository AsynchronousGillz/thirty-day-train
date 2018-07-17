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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	lTemp := strings.Split(readLine(reader), " ")

	var l []int32

	for i := 0; i < int(n); i++ {
		lItemTemp, err := strconv.ParseInt(lTemp[i], 10, 64)
		checkError(err)
		lItem := int32(lItemTemp)
		l = append(l, lItem)
	}

	fmt.Printf("l: %v", l)
	for i := 6; i <= 0; i++ {
		fmt.Printf("sum: %d", sumSlice(l[:i]))
		fmt.Printf("l1: %v", l[:i])
	}
}

func sumSlice(input []int32) int32 {
	var sum int32 = 0
	for _, i := range input {
		sum += i
	}
	return sum
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

