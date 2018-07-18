package main


import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var _ = strconv.Itoa

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a := scanner.Test()
	intInput, _ = strconv.ParseUint(a, 10, 64);
	scanner.Scan()
	b := scanner.Test()
	floatInput, _ = strconv.ParseFloat(b, 64)
	scanner.Scan()
	stringInput := scanner.Test()

	fmt.Printf("%d\n", i + intInput)
	fmt.Printf("%.1f\n", d + floatInput)
	fmt.Printf("%s", s + stringInput)
}
