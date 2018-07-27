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
    ret := 0
    prev := 0
    for i := n; i > 0; i = i/2 {
        remainder := i % 2
        if remainder == 1 {
            ret += 1
        } else {
            if ret > prev {
                prev = ret
            }
            ret = 0
        }
    }
    if ret > prev {
        prev = ret
    }
    fmt.Println(prev)
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
