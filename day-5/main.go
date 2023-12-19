package main

import (
	"fmt"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

func main() {
	scanner := helpers.NewScanner("input.txt")
	defer scanner.Close()
	part1, part2 := solve(scanner)
	fmt.Printf("part1 location number: %v\n", part1)
	fmt.Printf("part2 location number: %v\n", part2)
}
