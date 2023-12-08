package part1

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Point struct {
	X, Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type Coordinates struct {
	Start, End     Point
	AdjacentPoints []*Point
}

func (c *Coordinates) String() string {
	return fmt.Sprintf("(start: %s, end: %s)", c.Start.String(), c.End.String())
}

func (c *Coordinates) findAdjacentPoints() {
	// this accepts points that are outside the grid but ¯\_(ツ)_/¯
	for x := c.Start.X - 1; x <= c.End.X+1; x++ {
		for y := c.Start.Y - 1; y <= c.End.Y+1; y++ {
			switch {
			case x < 0, y < 0:
				continue
			case y == c.Start.Y && x >= c.Start.X && x <= c.End.X:
				continue
			default:
				c.AdjacentPoints = append(c.AdjacentPoints, &Point{x, y})
			}
		}
	}
}

type SchematicNumber struct {
	Number      int
	Coordinates Coordinates
}

func (s *SchematicNumber) String() string {
	return fmt.Sprintf("number: %d, coordinates: %s", s.Number, s.Coordinates.String())
}

func (s *SchematicNumber) findAdjacentPoints() {
	s.Coordinates.findAdjacentPoints()
}

type Schematics struct {
	Symbols        map[Point]string
	Numbers        []*SchematicNumber
	PartNumbers    []*SchematicNumber
	NonPartNumbers []*SchematicNumber
}

func (s *Schematics) parseLine(line string, lineNumber int) {
	pattern := `([^\w\s.])|(\d+)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		if match[2] != -1 {
			point := Point{match[2], lineNumber}
			s.Symbols[point] = line[match[2]:match[3]]
		}
		if match[4] != -1 {
			sNumber := &SchematicNumber{
				parseNumber(line, match[4], match[5]),
				Coordinates{Point{match[4], lineNumber}, Point{match[5] - 1, lineNumber}, nil},
			}
			sNumber.findAdjacentPoints()
			s.Numbers = append(s.Numbers, sNumber)
		}
	}
}

func parseNumber(line string, start, end int) int {
	number, _ := strconv.Atoi(line[start:end])
	return number
}

func ParseSchematics(scanner *bufio.Scanner) *Schematics {
	s := &Schematics{
		Symbols: make(map[Point]string),
	}
	lineNumber := 0
	for scanner.Scan() {
		s.parseLine(scanner.Text(), lineNumber)
		lineNumber++
	}

numberLoop:
	for _, number := range s.Numbers {
		for _, adjacentPoint := range number.Coordinates.AdjacentPoints {
			_, exists := s.Symbols[*adjacentPoint]
			if exists {
				s.PartNumbers = append(s.PartNumbers, number)
				continue numberLoop
			}
		}
		s.NonPartNumbers = append(s.NonPartNumbers, number)
	}
	return s
}
