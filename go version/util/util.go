package util

import (
	"os"
	"strings"
	"time"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadLines(path string) []string {
	asString := ReadFile(path)
	lines := strings.Split(strings.ReplaceAll(asString, "\r\n", "\n"), "\n")
	return lines
}

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	Check(err)
	return string(data)
}

func TestFunc(fn func()) time.Duration {
	start := time.Now()
	fn()
	end := time.Now()
	return end.Sub(start)
}

func Transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
