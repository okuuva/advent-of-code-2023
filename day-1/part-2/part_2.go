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

func getDigits(s string) []string {
	// FIXME: this does not work with overlapping numbers like oneight
	// should instead find the index for all numbers in the string and replace first and last spelled
	// numbers with their digit equivalent
	pattern := `\d|one|two|three|four|five|six|seven|eight|nine`
	re := regexp.MustCompile(pattern)
	return re.FindAllString(s, -1)
}

func LineToNumber(s string) int {
	digits := getDigits(s)
	if digits == nil {
		return 0
	}
	numberString := strings.Builder{}
	for _, digit := range [2]string{digits[0], digits[len(digits)-1]} {
		// thanks copilot for this switch statement
		// ...though without it I might've spent a little more time to figure a smarter way to do this
		switch digit {
		case "one", "1":
			numberString.WriteString("1")
		case "two", "2":
			numberString.WriteString("2")
		case "three", "3":
			numberString.WriteString("3")
		case "four", "4":
			numberString.WriteString("4")
		case "five", "5":
			numberString.WriteString("5")
		case "six", "6":
			numberString.WriteString("6")
		case "seven", "7":
			numberString.WriteString("7")
		case "eight", "8":
			numberString.WriteString("8")
		case "nine", "9":
			numberString.WriteString("9")
		}
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
