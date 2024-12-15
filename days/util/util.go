package util

import (
	"os"
	"strings"
	"time"
)

// useful functions
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

func Min[T int | uint | uint8](a, b T) T {
	if a > b {
		return b
	}
	return a
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

// directions

type Position struct {
	X, Y int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) OppositeDirection() Direction {
	return OppositeDirection(d)
}

func OppositeDirection(d Direction) Direction {
	switch d {
	case Up:
		return Down
	case Right:
		return Left
	case Down:
		return Up
	case Left:
		return Right
	default:
		panic("This Opposite Direction is not handled")
	}
}

func (p *Position) Move(d Direction) {
	switch d {
	case Up:
		p.Y--
	case Right:
		p.X++
	case Down:
		p.Y++
	case Left:
		p.X--
	}
}

func (p *Position) MoveAndCopy(d Direction) Position {
	cp := *p
	switch d {
	case Up:
		cp.Y--
	case Right:
		cp.X++
	case Down:
		cp.Y++
	case Left:
		cp.X--
	}
	return cp
}

func (p *Position) InBounds(width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}
