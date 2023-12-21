package main

// So, the first step is to put the hands in order of strength:
//
// 32T3K is the only one pair and the other hands are all a stronger type, so it gets rank 1.
// KK677 and KTJJT are both two pair. Their first cards both have the same label, but the second card of
//  KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
// T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card, so it gets rank 5 and T55J5 gets
//  rank 4.
//
// Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each
//  hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this
//  example are 6440.
import (
	"reflect"
	"testing"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

var expectedHands = hands{
	{hand: "32T3K", bid: 765, rank: 1, _type: onePair},
	{hand: "T55J5", bid: 684, rank: 4, _type: threeOfAKind},
	{hand: "KK677", bid: 28, rank: 3, _type: twoPairs},
	{hand: "KTJJT", bid: 220, rank: 2, _type: twoPairs},
	{hand: "QQQJA", bid: 483, rank: 5, _type: threeOfAKind},
}
var parsedHands = resetHands(&expectedHands)

var expectedExtraHands = hands{
	{hand: "", bid: 10, rank: 7, _type: fiveOfAKind},
	{hand: "", bid: 10, rank: 6, _type: fourOfAKind},
	{hand: "", bid: 10, rank: 5, _type: fullHouse},
	{hand: "", bid: 10, rank: 4, _type: threeOfAKind},
	{hand: "", bid: 10, rank: 3, _type: twoPairs},
	{hand: "", bid: 10, rank: 2, _type: onePair},
	{hand: "", bid: 10, rank: 1, _type: highCard},
}
var parsedExtraHands = resetHands(&expectedExtraHands)

var part1ResultMap = [][]hands{
	{expectedHands, parsedHands},
	{expectedExtraHands, parsedExtraHands},
}

const part1Solution = 6440

const part2Solution = 0

func TestParseInput(t *testing.T) {
	expected := &parsedHands
	s := helpers.NewScanner("testInput.txt")
	got := parseInput(s)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("parseInput: expected '%v', got '%v'", expected, got)
	}
}

func TestPart1Results(t *testing.T) {
	for _, hands := range part1ResultMap {
		expected, got := hands[0], hands[1]
		rankHands(&got)
		if !reflect.DeepEqual(&expected, &got) {
			t.Errorf("part1 results: expected '%v', got '%v'", expected, got)
		}
	}
}

func TestCountWinnings(t *testing.T) {
	expected := 70 + 60 + 50 + 40 + 30 + 20 + 10
	got := countWinnings(&expectedExtraHands)
	if got != expected {
		t.Errorf("countWinnings: expected '%v', got '%v'", expected, got)
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
