package main

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	input    string
	possible bool
	sets     []*Set
}{
	{
		input:    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		possible: true,
		sets: []*Set{
			{blue: 3, red: 4},
			{red: 1, green: 2, blue: 6},
			{green: 2},
		},
	},
	{
		input:    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		possible: true,
		sets: []*Set{
			{blue: 1, green: 2},
			{green: 3, blue: 4, red: 1},
			{green: 1, blue: 1},
		},
	},
	{
		input:    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		possible: false,
		sets: []*Set{
			{green: 8, blue: 6, red: 20},
			{blue: 5, red: 4, green: 13},
			{green: 5, red: 1},
		},
	},
	{
		input:    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		possible: false,
		sets: []*Set{
			{green: 1, red: 3, blue: 6},
			{green: 3, red: 6},
			{green: 3, blue: 15, red: 14},
		},
	},
	{
		input:    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		possible: true,
		sets: []*Set{
			{red: 6, blue: 1, green: 3},
			{blue: 2, red: 1, green: 2},
		},
	},
}

func TestIsPossible(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s should give %v", tc.input, tc.possible), func(t *testing.T) {
			game := parseLine(tc.input)
			got := game.isPossible()
			if got != tc.possible {
				t.Fatalf("game.isPossible(%q)\n got: %v, want: %v", tc.input, got, tc.possible)
			}
		})
	}
}

func TestParseSets(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Checking sets for %s", tc.input), func(t *testing.T) {
			game := parseLine(tc.input)
			if !reflect.DeepEqual(game.sets, tc.sets) {
				t.Fatalf("parseSets(%q)\n got: %v, want: %v", tc.input, game.sets, tc.sets)
			}
		})
	}
}

func TestSum(t *testing.T) {
	sum := 0
	expected := 0
	for _, tc := range testCases {
		game := parseLine(tc.input)
		if game.isPossible() {
			sum += game.id
		}
		if tc.possible {
			expected += game.id
		}
	}
	if sum != expected {
		t.Fatalf("sum of test cases should be %d, got %d", expected, sum)
	}
}
