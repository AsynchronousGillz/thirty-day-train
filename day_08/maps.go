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
    m := make(map[string]int64)

    for i := 0; i < int(n); i++ {
        arrTemp := strings.Split(readLine(reader), " ")
        tempItem, err := strconv.ParseInt(arrTemp[1], 10, 64)
        checkError(err)
        m[arrTemp[0]] = tempItem
    }

    for {
        arrTemp := readLine(reader)
        if arrTemp == "" {
            break
        }
        if _, ok := m[arrTemp]; ok {
            fmt.Printf("%s=%d\n", arrTemp, m[arrTemp])
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
