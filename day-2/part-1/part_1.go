package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var condition = struct {
	red, blue, green int
}{
	red:   12,
	green: 13,
	blue:  14,
}

type Set struct {
	red, blue, green int
}

func (s *Set) String() string {
	return fmt.Sprintf("{red: %d, blue: %d, green: %d}", s.red, s.blue, s.green)
}

type Game struct {
	id   int
	sets []*Set
}

func (g *Game) isPossible() bool {
	for _, set := range g.sets {
		if set.red > condition.red || set.blue > condition.blue || set.green > condition.green {
			return false
		}
	}
	return true
}

func parseGameId(s string) int {
	id, _ := strconv.Atoi(strings.Split(s, " ")[1])
	return id
}

func parseSets(s string) []*Set {
	sets := []*Set{}
	for _, setString := range strings.Split(s, ";") {
		set := parseSet(setString)
		sets = append(sets, set)
	}
	return sets
}

func parseSet(s string) *Set {
	pattern := `(\d+)\s(green|blue|red)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(s, -1)
	set := Set{}
	for _, match := range matches {
		number, err := strconv.Atoi(match[1])
		if err != nil {
			number = 0
		}
		switch match[2] {
		case "red":
			set.red = number
		case "green":
			set.green = number
		case "blue":
			set.blue = number
		}

	}
	return &set
}

func parseLine(s string) *Game {
	line := strings.SplitN(s, ":", 2)
	return &Game{
		id:   parseGameId(line[0]),
		sets: parseSets(line[1]),
	}
}

func main() {
	filename := "input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := parseLine(scanner.Text())
		if game.isPossible() {
			sum += game.id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during file scan: %s", err)
	}

	fmt.Println(sum)
}
