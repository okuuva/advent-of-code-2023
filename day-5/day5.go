package main

import (
	"math"
	"regexp"
	"strings"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

type almanacMap = [][]int
type almanacMaps = map[string]almanacMap
type sets = [][]int

var mapNames = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

var numberRe = regexp.MustCompile(`\d+`)

func parseHeader(s []string) []int {
	return helpers.Map(numberRe.FindAllString(s[0], -1), helpers.Atoi)
}

func parseMapSection(s []string) (string, [][]int) {
	mapName := strings.Split(s[0], " ")[0]
	ranges := make([][]int, len(s)-1)

	for i, line := range s[1:] {
		numbers := numberRe.FindAllString(line, -1)
		ranges[i] = helpers.Map(numbers, helpers.Atoi)
	}

	return mapName, ranges
}

func readSection(s *helpers.Scanner) []string {
	section := []string{}
	for {
		line := s.Text()
		if len(line) == 0 {
			break
		}
		section = append(section, line)
		s.Scan()
	}
	return section
}

func parseInput(s *helpers.Scanner) ([]int, almanacMaps) {
	s.Scan()
	seeds := parseHeader(readSection(s))
	almanacMaps := make(almanacMaps)
	for s.Scan() {
		mapName, ranges := parseMapSection(readSection(s))
		almanacMaps[mapName] = ranges
	}
	return seeds, almanacMaps
}

func traceMaps(s *helpers.Scanner) sets {
	seeds, almanacMaps := parseInput(s)
	sets := make(sets, len(seeds))
	for i, prev := range seeds {
		sets[i] = make([]int, len(mapNames)+1)
        sets[i][0] = prev
		for j, mapName := range mapNames {
			for _, _range := range almanacMaps[mapName] {
				destStart, sourceStart, length := _range[0], _range[1], _range[2]
				sourceEnd := sourceStart + length
				if prev >= sourceStart && prev <= sourceEnd {
					prev = destStart - sourceStart + prev
					break
				}
			}
			sets[i][j+1] = prev
		}
	}
	return sets
}

func solve(s *helpers.Scanner) (int, int) {
	sets := traceMaps(s)
	lowestLocation := math.MaxInt
	for _, set := range sets {
		location := set[len(set)-1]
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation, 0
}
