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

	for i := 0; i < num; i++ {
		word, _ := reader.ReadString('\n')
		word = strings.TrimSpace(word)
		fmt.Println(convert(word))
	}
}

func convert(word string) string {
	if len(word) <= 10 {
		return word
	}
	return fmt.Sprintf("%c%d%c", word[0], len(word)-2, word[len(word)-1])
}
