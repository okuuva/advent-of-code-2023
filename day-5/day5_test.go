package main

import (
	"reflect"
	"testing"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

var expectedSets = [][]int{
	{79, 81, 81, 81, 74, 78, 78, 82},
	{14, 14, 53, 49, 42, 42, 43, 43},
	{55, 57, 57, 53, 46, 82, 82, 86},
	{13, 13, 52, 41, 34, 34, 35, 35},
}

func TestTraceMaps(t *testing.T) {
	scanner := helpers.NewScanner("testInput.txt")
	defer scanner.Close()
	sets := traceMaps(scanner)
	for i, v := range sets {
		if !reflect.DeepEqual(v, expectedSets[i]) {
			t.Errorf("expected %v, got %v", expectedSets[i], v)
		}
	}
}

func TestSolve(t *testing.T) {
	scanner := helpers.NewScanner("testInput.txt")
	defer scanner.Close()
	part1, part2 := solve(scanner)
	if part1 != 35 {
		t.Errorf("part1: expected 35, got %d", part1)
	}
	if part2 != 0 {
		t.Errorf("part2: expected 0, got %d", part2)
	}
}
