package main
import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	fmt.Printf("Hello, World.\n%s", input)
}
