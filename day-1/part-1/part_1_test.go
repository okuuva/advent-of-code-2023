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
		input:    "1abc2",
		expected: 12,
	},
	{
		input:    "pqr3stu8vwx",
		expected: 38,
	},
	{
		input:    "a1b2c3d4e5f",
		expected: 15,
	},
	{
		input:    "treb7uchet",
		expected: 77,
	},
	{
		input:    "trebuchet",
		expected: 0,
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
