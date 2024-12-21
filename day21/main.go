package main

import (
    "fmt"
    "strconv"
)

var keypad = map[string][2]int{
	"7": {0, 0}, "8": {0, 1}, "9": {0, 2},
	"4": {1, 0}, "5": {1, 1}, "6": {1, 2},
	"1": {2, 0}, "2": {2, 1}, "3": {2, 2},
	"0": {3, 1}, "A": {3, 2},

	"^": {0, 1}, "B": {0, 2},
	"<": {1, 0}, "v": {1, 1}, ">": {1, 2},
}

var dirs = map[rune][2]int{
	'^': {-1, 0},
	'<': {0, -1},
	'v': {1, 0},
	'>': {0, 1},
}

type memoKey struct {
    code       string
    maxIter    int
    depth      int
}
var memo = make(map[memoKey]int)

func permutationsOfString(s string) []string {
	runes := []rune(s)
	var results []string
	seen := make(map[string]bool)

	var generate func(int)
	generate = func(start int) {
		if start == len(runes)-1 {
			perm := string(runes)
			if !seen[perm] {
				seen[perm] = true
				results = append(results, perm)
			}
			return
		}
		for i := start; i < len(runes); i++ {
			runes[start], runes[i] = runes[i], runes[start]
			generate(start + 1)
			runes[start], runes[i] = runes[i], runes[start] // backtrack
		}
	}

	generate(0)
	return results
}

func checkmoves(start, end, illegal [2]int) []string {
	dx := end[0] - start[0]
	dy := end[1] - start[1]

	movesNeeded := ""
	if dx < 0 {
		for i := 0; i < -dx; i++ {
			movesNeeded += "^"
		}
	}
	if dx > 0 {
		for i := 0; i < dx; i++ {
			movesNeeded += "v"
		}
	}
	if dy < 0 {
		for i := 0; i < -dy; i++ {
			movesNeeded += "<"
		}
	}
	if dy > 0 {
		for i := 0; i < dy; i++ {
			movesNeeded += ">"
		}
	}

	perms := permutationsOfString(movesNeeded)

	validMoves := []string{}
Outer:
	for _, perm := range perms {
		pos := start
		for _, step := range perm {
			d := dirs[step]
			pos = [2]int{pos[0] + d[0], pos[1] + d[1]}
			if pos == illegal {
				continue Outer
			}
		}
		validMoves = append(validMoves, perm+"B")
	}

	if len(validMoves) == 0 {
		return []string{"B"}
	}
	return validMoves
}

func findPath(code string, maxiterations, depth int) int {
	key := memoKey{code, maxiterations, depth}
	if val, ok := memo[key]; ok {
		return val
	}

	var cur, illegal [2]int
	if depth == 0 {
		cur = keypad["A"]
		illegal = [2]int{3, 0}
	} else {
		cur = keypad["B"]
		illegal = [2]int{0, 0}
	}

	totalLength := 0
	for _, ch := range code {
		next := keypad[string(ch)]
		allMoves := checkmoves(cur, next, illegal)

		if depth == maxiterations {
			totalLength += len(allMoves[0])
		} else {
			minVal := -1
			for _, moveStr := range allMoves {
				cost := findPath(moveStr, maxiterations, depth+1)
				if minVal == -1 || cost < minVal {
					minVal = cost
				}
			}
			totalLength += minVal
		}
		cur = next
	}

	memo[key] = totalLength
	return totalLength
}

func main() {
	rawdata := []string{
		// "480A",
		// "143A",
		// "983A",
		// "382A",
		// "974A",

		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}

	//part1 2 iterations
	result := 0
	for _, code := range rawdata {
		if len(code) < 2 {
			continue
		}
		numericPart := code[:len(code)-1]
		numeric, err := strconv.Atoi(numericPart)
		if err != nil {
			continue
		}
		result += findPath(code, 2, 0) * numeric
	}
	fmt.Println(result)

	//part2 25 iterations
	result = 0
	for _, code := range rawdata {
		if len(code) < 2 {
			continue
		}
		numericPart := code[:len(code)-1]
		numeric, err := strconv.Atoi(numericPart)
		if err != nil {
			continue
		}
		result += findPath(code, 25, 0) * numeric
	}
	fmt.Println(result)
}