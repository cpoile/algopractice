package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, _ := strconv.Atoi(numStr)

	fmt.Println(hasEvenFactor(num))
}

func hasEvenFactor(num int) string {
	for i := 2; i <= num/2; i++ {
		if num%2 == 0 && num%i == 0 {
			return "YES"
		}
	}
	return "NO"
}
