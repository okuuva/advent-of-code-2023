package main

import (
	"fmt"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

func part1() {
	scanner := helpers.NewScanner("input.txt")
	defer scanner.Close()

	sum := 0
	for scanner.Scan() {
		card := ParseCard(scanner.Text())
		sum += card.CountPoints()
	}
	fmt.Printf("part1 sum: %v\n", sum)
}

func main() {
	part1()
}
