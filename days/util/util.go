package util

import (
	"os"
	"strings"
	"time"
)

type Position struct {
	X, Y int
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadLines(path string) []string {
	asString := ReadFile(path)
	lines := strings.Split(CleanInput(asString), "\n")
	return lines
}

func CleanInput(input string) string {
	return strings.ReplaceAll(input, "\r\n", "\n")
}

func Map[T any, V any](slice []T, fn func(T) V) []V {
	newSlice := make([]V, len(slice))
	for i, element := range slice {
		newSlice[i] = fn(element)
	}
	return newSlice
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	newSlice := make([]T, 0, len(slice))
	for _, element := range slice {
		if fn(element) {
			newSlice = append(newSlice, element)
		}
	}
	return newSlice
}

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	Check(err)
	return CleanInput(string(data))
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
