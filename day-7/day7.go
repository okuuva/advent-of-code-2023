package main

import (
	"slices"
	"strings"

	"github.com/okuuva/advent-of-code-2023/helpers"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

var cardRanks = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 3,
	'3': 3,
	'2': 2,
}

type handType int

const (
	fiveOfAKind = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPairs
	onePair
	highCard
)

type hand struct {
	hand  string
	bid   int
	rank  int
	_type handType
}

func (h *hand) winnings() int {
	return h.bid * h.rank
}

func (h *hand) parseType() handType {
	values := maps.Values(lo.CountValues(strings.Split(h.hand, "")))
	valuesCount := len(values)
	valuesMax := lo.Max(values)
	switch valuesCount {
	case 1:
		h._type = fiveOfAKind
	case 2:
		if valuesMax == 4 {
			h._type = fourOfAKind
		} else {
			h._type = fullHouse
		}
	case 3:
		if valuesMax == 3 {
			h._type = threeOfAKind
		} else {
			h._type = twoPairs
		}
	case 4:
		h._type = onePair
	default:
		h._type = highCard
	}
	return h._type
}

func (h *hand) cmp(other *hand) int {
	otherRunes := []rune(other.hand)
	for i, char := range h.hand {
		switch {
		case cardRanks[char] > cardRanks[otherRunes[i]]:
			return 1
		case cardRanks[char] < cardRanks[otherRunes[i]]:
			return -1
		}
	}
	return 0
}

type hands = []*hand

func resetHands(h *hands) hands {
	// reset rank and _type
	var newHands = make(hands, len(*h))
	for i, hand := range *h {
		hand.rank = 0
		hand._type = -1
		newHands[i] = hand
	}
	return newHands
}

func parseInput(s *helpers.Scanner) *hands {
	hands := hands{}
	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		hands = append(hands, &hand{
			hand:  line[0],
			bid:   helpers.Atoi(line[1]),
			rank:  0,
			_type: -1,
		})
	}
	return &hands
}

func rankHands(h *hands) {
	handsByRank := make([]hands, 7)
	for _, hand := range *h {
		rank := hand.parseType()
		handsByRank[rank] = append(handsByRank[rank], hand)
	}
	rank := len(*h)
	for _, hands := range handsByRank {
		slices.SortFunc(hands, func(a, b *hand) int {return a.cmp(b)})
		slices.Reverse(hands)
		for _, hand := range hands {
			hand.rank = rank
			rank--
		}
	}
}

func countWinnings(h *hands) int {
	return lo.Reduce(*h, func(agg int, item *hand, _ int) int { return agg + item.winnings() }, 0)
}

func solve(s *helpers.Scanner) (int, int) {
	hands := parseInput(s)
	rankHands(hands)
	part1 := countWinnings(hands)
	part2 := 0
	return part1, part2
}
