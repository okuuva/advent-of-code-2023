package main

import (
	"fmt"
	"testing"
)

var testCasesPart1 = []struct {
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

var testCasesPart2 = []struct {
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
}

func TestLineToNumberPart1(t *testing.T) {
	for _, tc := range testCasesPart1 {
		t.Run(fmt.Sprintf("%s should give %d", tc.input, tc.expected), func(t *testing.T) {
			got := LineToNumber(tc.input)
			if got != tc.expected {
				t.Fatalf("LineToNumber(%q)\n got: %v, want: %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestSumPart1(t *testing.T) {
	sum := 0
	expected := 0
	for _, tc := range testCasesPart1 {
		sum += LineToNumber(tc.input)
		expected += tc.expected
	}
	if sum != expected {
		t.Fatalf("sum of test cases should be %d, got %d", expected, sum)
	}
}

func TestLineToNumberPart2(t *testing.T) {
	for _, tc := range testCasesPart2 {
		t.Run(fmt.Sprintf("%s should give %d", tc.input, tc.expected), func(t *testing.T) {
			got := LineToNumber(tc.input)
			if got != tc.expected {
				t.Fatalf("LineToNumber(%q)\n got: %v, want: %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestSumPart2(t *testing.T) {
	sum := 0
	expected := 0
	for _, tc := range testCasesPart2 {
		sum += LineToNumber(tc.input)
		expected += tc.expected
	}
	if sum != expected {
		t.Fatalf("sum of test cases should be %d, got %d", expected, sum)
	}
}
