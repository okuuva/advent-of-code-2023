package main

import (
	"bufio"
	p1 "day3/part1"
	"fmt"
	"os"
)

func main() {
	filename := "input.txt"

	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	schematics := p1.ParseSchematics(scanner)
	sum := 0

	for _, number := range schematics.PartNumbers {
		sum += number.Number
	}

	fmt.Println(sum)
}
