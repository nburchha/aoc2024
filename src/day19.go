package main

import (
	"fmt"
	"os"
	// "strconv"
	"strings"
	// "regexp"
)

const (
	FILE = "input/day19.txt"
	// FILE = "input/testinput19"
)



func parseTowels(lines []string) ([]string, []string) {
	var towels, patterns []string
	split := strings.Split(lines[0], ", ")
	for _, s := range split {
		towels = append(towels, s)
	}
	for _, s := range lines[2:] {
		patterns = append(patterns, s)
	}
	return towels, patterns
}

func possiblePattern(towels []string, pattern string, index int, memo map[int]int) int {
	if index == len(pattern) {
		return 1
	}
	if val, exists := memo[index]; exists {
		return val
	}
	count := 0
	for _, towel := range towels {
		if index+len(towel) <= len(pattern) && strings.HasPrefix(pattern[index:], towel) {
			count += possiblePattern(towels, pattern, index+len(towel), memo)
		}
	}
	memo[index] = count
	return count
}

func main() {
	input, _ := os.ReadFile(FILE)
	lines := strings.Split(string(input), "\n")
	towels, patterns := parseTowels(lines)
	// fmt.Println(towels, patterns)
	var possible int
	var count int
	for _, pattern := range patterns {
		if tmp := possiblePattern(towels, pattern, 0, make(map[int]int)); tmp > 0  {
			count += tmp
			possible++
		}
	}
	fmt.Println(possible, count)
}