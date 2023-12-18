package main

import (
	"fmt"

	"github.com/okuuva/advent-of-code-2023/helpers"
)

func main() {
	scanner := helpers.NewScanner("input.txt")
	defer scanner.Close()
	locationNumber, _ := solve(scanner)
	fmt.Printf("part1 location number: %v\n", locationNumber)
}
