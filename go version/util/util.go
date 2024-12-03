package util

import (
	"os"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadLines(path string) []string {
	data, err := os.ReadFile(path)
	Check(err)
	asString := string(data)
	lines := strings.Split(strings.ReplaceAll(asString, "\r\n", "\n"), "\n")
	return lines
}
