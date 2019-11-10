package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	input.Scan()
	n, _ := strconv.Atoi(input.Text())
	input.Scan()
	k, _ := strconv.Atoi(input.Text())

	count := 0
	var ksScore int
	for place := 1; place <= n; place, count = place+1, count+1 {
		input.Scan()
		score, _ := strconv.Atoi(input.Text())
		if score < 1 {
			break
		}
		if place == k {
			ksScore = score
		}
		if place > k && score < ksScore {
			break
		}
	}

	fmt.Println(count)
}
