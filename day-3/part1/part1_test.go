package part1

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

var testCase = struct {
	schematic      string
	partNumbers    []int
	nonPartNumbers []int
	expectedSum    int
}{
	schematic: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
	partNumbers:    []int{467, 35, 633, 617, 592, 755, 664, 598},
	nonPartNumbers: []int{114, 58},
	expectedSum:    4361,
}

func TestPartNumberCount(t *testing.T) {
	s := ParseSchematics(bufio.NewScanner(bytes.NewReader([]byte(testCase.schematic))))
	got, expected := len(s.PartNumbers), len(testCase.partNumbers)
	if got != expected {
		t.Fatalf("expected %d part numbers, got %d", expected, got)
	}
}

func TestNonPartNumberCount(t *testing.T) {
	s := ParseSchematics(bufio.NewScanner(bytes.NewReader([]byte(testCase.schematic))))
	got, expected := len(s.NonPartNumbers), len(testCase.nonPartNumbers)
	if got != expected {
		t.Fatalf("expected %d non part numbers, got %d", expected, got)
	}
}

func TestPartNumbers(t *testing.T) {
	s := ParseSchematics(bufio.NewScanner(bytes.NewReader([]byte(testCase.schematic))))
	for _, number := range testCase.partNumbers {
		t.Run(fmt.Sprintf("%v should be part number", number), func(t *testing.T) {
			for _, comparison := range s.PartNumbers {
				if comparison.Number == number {
					return
				}
			}
			t.Fatalf("%v should be part number", number)
		})
	}
}

func TestNonPartNumbers(t *testing.T) {
	s := ParseSchematics(bufio.NewScanner(bytes.NewReader([]byte(testCase.schematic))))
	for _, number := range testCase.nonPartNumbers {
		t.Run(fmt.Sprintf("%v should be non part number", number), func(t *testing.T) {
			for _, comparison := range s.NonPartNumbers {
				if comparison.Number == number {
					return
				}
			}
			t.Fatalf("%v should be non part number", number)
		})
	}
}

func TestSum(t *testing.T) {
	sum := 0
	s := ParseSchematics(bufio.NewScanner(bytes.NewReader([]byte(testCase.schematic))))
	for _, number := range s.PartNumbers {
		sum += number.Number
	}
	expected := testCase.expectedSum
	if sum != expected {
		t.Fatalf("sum of test cases should be %d, got %d", expected, sum)
	}
}
