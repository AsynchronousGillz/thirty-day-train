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
	m := make(map[string]int)

	for i := 0; i < int(n); i++ {
		arrTemp := strings.Split(readLine(reader), " ")
		tempItem, err := strconv.ParseFloat(arrTemp[1], 64)
		checkError(err)
		m[arrTemp[0]] = float64(tempItem)
	}

	for i := 0; i < int(n); i++ {
		arrTemp := readLine(reader)
		m[arrTemp[0]] = arrItem
		if i, ok := m[arrItem[i]]; ok {
			fmt.Printl(m[arrItem[i]])
		} else {
			fmt.Println("Not found")
		}
	}
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

