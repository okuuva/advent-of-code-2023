package main

import (
	"reflect"
	"testing"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

var part1TestRaces = []race{
	{7, 9}, {15, 40}, {30, 200},
}
var part1Results = []int{4, 8, 9}
var part1Solution = helpers.Product(part1Results)

var part2TestRace = race{71530, 940200}
var part2Solution = 71503

func TestParseInput(t *testing.T) {
	s := helpers.NewScanner("testInput.txt")
	got := parseInput(s)
	if !reflect.DeepEqual(got, part1TestRaces) {
		t.Errorf("parseInput: expected '%v', got '%v'", part1TestRaces, got)
	}
}

func TestFixParsedInput(t *testing.T) {
	got := fixParsedInput(part1TestRaces)
	if !reflect.DeepEqual(got, part2TestRace) {
		t.Errorf("parseInputPart2: expected '%v', got '%v'", part2TestRace, got)
	}
}

func TestPart1Results(t *testing.T) {
	got := numbersOfWinningOptions(part1TestRaces)
	if !reflect.DeepEqual(got, part1Results) {
		t.Errorf("part1 results: expected '%v', got '%v'", part1Results, got)
	}
}

func TestSolve(t *testing.T) {
	s := helpers.NewScanner("testInput.txt")
	part1, part2 := solve(s)
	if part1 != part1Solution {
		t.Errorf("part1 solution: expected '%v', got '%v'", part1Solution, part1)
	}
	if part2 != part2Solution {
		t.Errorf("part2 solution: expected '%v', got '%v'", part2Solution, part2)
	}
}
