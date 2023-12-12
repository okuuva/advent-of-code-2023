package main

import (
	"testing"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

var testCases = []struct {
	input          string
	id             int
	numbers        []int
	winningNumbers []int
	points         int
}{
	{
		input:          "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		id:             1,
		numbers:        []int{41, 48, 83, 86, 17},
		winningNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
		points:         8,
	},
	{
		input:          "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		id:             2,
		numbers:        []int{13, 32, 20, 16, 61},
		winningNumbers: []int{61, 30, 68, 82, 17, 32, 24, 19},
		points:         2,
	},
	{
		input:          "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		id:             3,
		numbers:        []int{1, 21, 53, 59, 44},
		winningNumbers: []int{69, 82, 63, 72, 16, 21, 14, 1},
		points:         2,
	},
	{
		input:          "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		id:             4,
		numbers:        []int{41, 92, 73, 84, 69},
		winningNumbers: []int{59, 84, 76, 51, 58, 5, 54, 83},
		points:         1,
	},
	{
		input:          "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		id:             5,
		numbers:        []int{87, 83, 26, 28, 32},
		winningNumbers: []int{88, 30, 70, 12, 93, 22, 82, 36},
		points:         0,
	},
	{
		input:          "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		id:             6,
		numbers:        []int{31, 18, 13, 56, 72},
		winningNumbers: []int{74, 77, 10, 23, 35, 67, 36, 11},
		points:         0,
	},
}

func TestParseCard(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			card := ParseCard(tc.input)
			if card.Id != tc.id {
				t.Errorf("expected id %d, got %d", tc.id, card.Id)
			}
			for i, number := range card.Numbers {
				if number != tc.numbers[i] {
					t.Errorf("expected number %d, got %d", tc.numbers[i], number)
				}
			}
			for i, number := range card.WinningNumbers {
				if number != tc.winningNumbers[i] {
					t.Errorf("expected winning number %d, got %d", tc.winningNumbers[i], number)
				}
			}
		})
	}
}

func TestCountPoints(t *testing.T) {
	sum := 0
	expected := 0
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			card := ParseCard(tc.input)
			got := card.CountPoints()
			sum += got
			expected += tc.points
			if got != tc.points {
				t.Fatalf("expected %d, got %d", tc.points, got)
			}
		})
	}
	if sum != expected {
		t.Fatalf("expected sum %d, got %d", expected, sum)
	}
}

func TestSolve(t *testing.T) {
	scanner := helpers.NewScanner("testInput.txt")
	defer scanner.Close()
	sum, totalNumberOfCards := Solve(scanner)
	if sum != 13 {
		t.Errorf("expected sum 13, got %d", sum)
	}
	if totalNumberOfCards != 30 {
		t.Errorf("expected totalNumberOfCards 30, got %d", totalNumberOfCards)
	}
}
