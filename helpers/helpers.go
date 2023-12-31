// Package helpers provides helper functions for Advent of Code.
package helpers

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/samber/lo"
)

// Scanner is a wrapper around bufio.Scanner that also has a Close method.
type Scanner struct {
	*bufio.Scanner
	Close func() error
}

// NewScanner returns a new Scanner for the given filepath. If opening path fails log.Fatalf is called.
func NewScanner(filepath string) *Scanner {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)

	return &Scanner{
		Scanner: scanner,
		Close:   file.Close,
	}
}

var parseIdRe = regexp.MustCompile(`\d+`)

// ParseId parses a numerical id from a simple string, such as "Game 1" or "Card 2".
func ParseId(s string) int {
	id := parseIdRe.FindString(s)
	numericalId, err := strconv.Atoi(id)
	if id == "" || err != nil {
		log.Fatalf("failed to parse id from string '%s': %s", s, err)
	}
	return numericalId
}

// Atoi is a wrapper around strconv.Atoi that calls log.Fatalf on error.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("failed to parse string '%s' to int: %s", s, err)
	}
	return i
}

// AtoiSlice is a wrapper around Atoi that calls it for each string in the given slice.
// Deprecated: use Map(s, Atoi) instead.
func AtoiSlice(s []string) []int {
	var numbers = make([]int, len(s))
	for i, v := range s {
		numbers[i] = Atoi(v)
	}
	return numbers
}

// Map applies the given function to each element in the given slice and returns a new slice with the results.
func Map[T, V any](slice []T, fn func(T) V) []V {
	result := make([]V, len(slice))
	for i, t := range slice {
		result[i] = fn(t)
	}
	return result
}

// MakeRange creates a slice containing integers from start to end inclusive. Returns an empty slice if start == end.
func MakeRange(start, end int) []int {
	switch {
	case end < start:
		s := MakeRange(end, start)
		slices.Reverse(s)
		return s
	case end == start:
		return []int{}
	}

	s := make([]int, end-start+1)
	for i, j := 0, start; j <= end; i, j = i+1, j+1 {
		s[i] = j
	}
	return s
}

// Product calculates the product of a slice of integers by multiplying all the elements together.
func Product(s []int) int {
	return lo.Reduce(s, func(agg, item, _ int) int { return agg * item }, 1)
}
