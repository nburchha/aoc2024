package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
)

func parseInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func parseMatches(line string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(line, -1)

	x,y := 0,0
	var count int
	for _, match := range matches {
		if (len(match) == 3) {
			fmt.Sscanf(match[1], "%d", &x)
			fmt.Sscanf(match[2], "%d", &y)
		}
		if (x != 0 && y != 0) {
			count += (x * y)
			x, y = 0, 0
		}
	}
	return count
}

func parseLine(line string) int {
	// Regex for `do()`, `dont()`, and `mul(x,y)`
	re := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\))`)

	// Start in active mode
	active := true
	count := 0

	// Find all matches
	matches := re.FindAllStringSubmatch(line, -1)

	// var x, y int
	for _, match := range matches {
		// Toggle state on `do()` or `dont()`
		if match[0] == "do()" {
			active = true
			continue
		} else if match[0] == "don't()" {
			active = false
			continue
		}
		if (!active) {
			continue
		}
		fmt.Println(match)
		x, errX := strconv.Atoi(match[2])
		y, errY := strconv.Atoi(match[3])
		if errX == nil || errY == nil {
			count += x * y
			fmt.Println("x:",x, "y", y, "count", count)
			x,y = 0,0
		}
		// fmt.Sscanf(match[2], "%d", &x)
		// fmt.Sscanf(match[3], "%d", &y)
	}
	return count
}

func part2(lines []string) (int) {
	count := 0
	for line := range lines {
		count += parseLine(lines[line])
	}
	return count
}

func part1(lines []string) int {
	var count int
	for line := range lines {
		count += parseMatches(lines[line])
	}
	return count
}

func main() {
	lines, err := parseInput("input/day03.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("part1:",part1(lines))
	fmt.Println("part2:",part2(lines))
}