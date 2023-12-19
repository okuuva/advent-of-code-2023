package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

type race = []int
type races = []race

var numberRe = regexp.MustCompile(`\d+`)

func parseInput(s *helpers.Scanner) races {
	lines := make([][]int, 2)
	for i := 0; s.Scan(); i++ {
		lines[i] = helpers.Map(numberRe.FindAllString(s.Text(), -1), helpers.Atoi)
	}
	numberOfRaces := len(lines[0])
	races := make(races, numberOfRaces)
	for range lines {
		for i := 0; i < numberOfRaces; i++ {
			races[i] = race{lines[0][i], lines[1][i]}
		}
	}
	return races
}

func fixParsedInput(races races) race {
	time, distance := strings.Builder{}, strings.Builder{}
	for _, r := range races {
		time.WriteString(strconv.Itoa(r[0]))
		distance.WriteString(strconv.Itoa(r[1]))
	}
	return race{helpers.Atoi(time.String()), helpers.Atoi(distance.String())}
}

func numberOfWinningOptions(r race) int {
	options := 0
	totalTime, targetDistance := r[0], r[1]
	for chargeTime := 0; chargeTime <= totalTime; chargeTime++ {
		speed := chargeTime
		if speed*(totalTime-chargeTime) > targetDistance {
			options++
		}
	}
	return options
}

func numbersOfWinningOptions(races races) []int {
	s := make([]int, len(races))
	for i, r := range races {
		s[i] = numberOfWinningOptions(r)
	}
	return s
}

func solve(s *helpers.Scanner) (int, int) {
	races := parseInput(s)
	part1 := helpers.Product(numbersOfWinningOptions(races))
	part2 := numberOfWinningOptions(fixParsedInput(races))
	return part1, part2
}
