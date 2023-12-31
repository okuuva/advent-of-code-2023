package main

import (
	"regexp"
	"sort"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

type Card struct {
	Id             int
	Numbers        []int
	WinningNumbers []int
	Points         int
	Matches        int
}

func (c *Card) CountPoints() int {
	numbers := make([]int, len(c.Numbers))
	copy(numbers, c.Numbers)
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })
	winningNumbers := make([]int, len(c.WinningNumbers))
	copy(winningNumbers, c.WinningNumbers)
	sort.Slice(winningNumbers, func(i, j int) bool { return winningNumbers[i] < winningNumbers[j] })

	points := 0
	for _, number := range numbers {
		for _, winningNumber := range winningNumbers {
			switch {
			case number > winningNumber:
				continue
			case number == winningNumber:
				c.Matches += 1
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
			break
		}
	}
	c.Points = points
	return points
}

var parseNumbersRe = regexp.MustCompile(`:([\d\s]+)\|([\d\s]+)`)
var separateNumbersRe = regexp.MustCompile(`\d+`)

func ParseCard(s string) *Card {
	id := helpers.ParseId(s)
	allNumbers := parseNumbersRe.FindStringSubmatch(s)
	parsedNumbers := separateNumbersRe.FindAllString(allNumbers[1], -1)
	parsedWinningNumbers := separateNumbersRe.FindAllString(allNumbers[2], -1)
	return &Card{
		Id:             id,
		Numbers:        helpers.AtoiSlice(parsedNumbers),
		WinningNumbers: helpers.AtoiSlice(parsedWinningNumbers),
	}
}

func Solve(scanner *helpers.Scanner) (int, int) {
	pointSum := 0
	var numberOfCards = make(map[int]int)
	totalNumberOfCards := 0
	for scanner.Scan() {
		card := ParseCard(scanner.Text())
		pointSum += card.CountPoints()

		numberOfCards[card.Id] += 1
		totalNumberOfCards += numberOfCards[card.Id]
		for i := card.Id + 1; i <= card.Id+card.Matches; i++ {
			numberOfCards[i] += numberOfCards[card.Id]
		}
	}
	return pointSum, totalNumberOfCards
}
