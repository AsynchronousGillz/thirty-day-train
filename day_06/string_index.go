package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var T int
	fmt.Scan(&T)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < T; i++ {
		word, _ := reader.ReadString('\n')
		word = strings.Replace(word, "\n", "", -1)
		for index := 0; index < len(word); index += 2 {
			fmt.Printf("%c", word[index])
		}
		fmt.Printf(" ")
		for index := 1; index < len(word); index += 2 {
			fmt.Printf("%c", word[index])
		}
		fmt.Println()
	}
}
