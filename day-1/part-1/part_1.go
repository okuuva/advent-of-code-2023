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
	pattern := `\d`
	re := regexp.MustCompile(pattern)
	return re.FindAllString(s, -1)
}

func LineToNumber(s string) int {
	digits := getDigits(s)
	if digits == nil {
		return 0
	}
	numberString := strings.Builder{}
    numberString.WriteString(digits[0])
    numberString.WriteString(digits[len(digits)-1])
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
