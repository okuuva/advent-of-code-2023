package main

import (
	"fmt"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

func main() {
	scanner := helpers.NewScanner("input.txt")
	defer scanner.Close()
	sum, totalNumberOfCards := Solve(scanner)
	fmt.Printf("part1 sum: %v\n", sum)
	fmt.Printf("part2 number of cards: %v\n", totalNumberOfCards)
}
