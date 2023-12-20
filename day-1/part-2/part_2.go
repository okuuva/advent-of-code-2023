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

// This wouldn't need a capture group but selecting results is easier when the output is in the same format
var firstNumberRe = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
var lastNumberRe = regexp.MustCompile(`.*(\d|one|two|three|four|five|six|seven|eight|nine).*$`)
var finders = []*regexp.Regexp{firstNumberRe, lastNumberRe}

var numberToDigit = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func LineToNumber(s string) int {
	numberString := strings.Builder{}
	for _, finder := range finders {
		result := finder.FindStringSubmatch(s)
		half := ""
		if result == nil {
			half = "0"
		} else {
			half = result[1]
		}
		digit, exists := numberToDigit[half]
		if !exists {
			digit = half
		}
		numberString.WriteString(digit)
	}
	number, _ := strconv.Atoi(numberString.String())
	return number
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
		sum += LineToNumber(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during file scan: %s", err)
	}

	fmt.Printf("sum: %d\n", sum)
}
