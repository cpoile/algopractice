package ocr

import (
	"bytes"
	"strings"
)

var numbers = [][]byte{
	[]byte(" _ | ||_|   "),
	[]byte("     |  |   "),
	[]byte(" _  _||_    "),
	[]byte(" _  _| _|   "),
	[]byte("   |_|  |   "),
	[]byte(" _ |_  _|   "),
	[]byte(" _ |_ |_|   "),
	[]byte(" _   |  |   "),
	[]byte(" _ |_||_|   "),
	[]byte(" _ |_| _|   ")}

func stringToBytes(s string) [][]byte {
	lines := strings.Split(s, "\n")
	ret := make([][]byte, len(lines))
	for i, l := range lines {
		ret[i] = []byte(l)
	}
	return ret
}

func Recognize(in string) []string {
	ins := stringToBytes(in[1:])
	var ret []string
	for i := 0; i < len(ins); i += 4 {
		ret = append(ret, recognizeLine(ins[i:i+4]))
	}

	return ret
}

func recognizeLine(in [][]byte) string {
	var ret []byte
	for i := 0; i < len(in[0]); i += 3 {
		var char []byte
		for j := 0; j < 4; j++ {
			char = append(char, in[j][i:i+3]...)
		}
		ret = append(ret, recognizeDigit(char))
	}
	return string(ret)
}

func recognizeDigit(in []byte) byte {
	for i, n := range numbers {
		if bytes.Equal(in, n) {
			return '0' + byte(i)
		}
	}
	return '?'
}

// also see: https://exercism.io/tracks/go/exercises/ocr-numbers/solutions/65b200a82276428f9b820bbf569d3786
