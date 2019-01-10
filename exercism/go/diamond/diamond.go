package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 65 || char > 90 {
		return "", errors.New("out of range")
	}
	diff := int(char - 'A')
	width := diff*2 + 1
	var ret string
	for i := diff; i >= -diff; i-- {
		if abs(i) == diff {
			ret += strings.Repeat(" ", abs(i)) + string(int(char)-abs(i)) + strings.Repeat(" ", abs(i))
		} else {
			ret += strings.Repeat(" ", abs(i)) + string(int(char)-abs(i)) +
				strings.Repeat(" ", width-2-abs(i)*2) +
				string(int(char)-abs(i)) + strings.Repeat(" ", abs(i))
		}
		ret += "\n"
	}
	return ret, nil
}

func abs(char int) int {
	if char < 0 {
		return -char
	}
	return char
}

// really like: https://exercism.io/tracks/go/exercises/diamond/solutions/4768dde9406e43c89d76ec40296a44b2
//
//func Gen(b byte) (string, error) {
//	if b < 'A' || b > 'Z' {
//		return "", errors.New("out of range")
//	}
//	x := int(b - 'A')
//	rows := make([][]byte, 2*x+1)
//	for i, j := x, 0; i >= 0; i, j = i-1, j+1 {
//		line := bytes.Repeat([]byte{' '}, 2*x+2)
//		line[2*x+1] = '\n'
//		char := 'A' + byte(j)
//		line[i], line[2*x-i] = char, char
//		rows[j], rows[2*x-j] = line, line
//	}
//	return string(bytes.Join(rows, nil)), nil
//}
