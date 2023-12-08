package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	input    string
	expected int
}{
	{
		input:    "two1nine",
		expected: 29,
	},
	{
		input:    "eightwothree",
		expected: 83,
	},
	{
		input:    "abcone2threexyz",
		expected: 13,
	},
	{
		input:    "xtwone3four",
		expected: 24,
	},
	{
		input:    "4nineeightseven2",
		expected: 42,
	},
	{
		input:    "zoneight234",
		expected: 14,
	},
	{
		input:    "7pqrstsixteen",
		expected: 76,
	},
	{
		input:    "two2geight",
		expected: 28,
	},
	{
		input:    "xxxxxxxxxxxxxxxxxxonexxxxxxxxxxxxxx",
		expected: 11,
	},
	{
		input:    "six3one4sixeighttwosbkqdjfhfroneights",
		expected: 68,
	},
	{
		input:    "oneeight",
		expected: 18,
	},
	{
		input:    "oneight",
		expected: 18,
	},
}

func TestLineToNumber(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s should give %d", tc.input, tc.expected), func(t *testing.T) {
			got := LineToNumber(tc.input)
			if got != tc.expected {
				t.Fatalf("LineToNumber(%q)\n got: %v, want: %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestSum(t *testing.T) {
	sum := 0
	expected := 0
	for _, tc := range testCases {
		sum += LineToNumber(tc.input)
		expected += tc.expected
	}
	if sum != expected {
		t.Fatalf("sum of test cases should be %d, got %d", expected, sum)
	}
}
