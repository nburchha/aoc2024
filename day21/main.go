package main

import (
	"fmt"
	"strconv"
)

var keyInit = Point{ row: 3, col: 2}
var RobInit = Point{ row: 0, col: 1}

var keyMap = map[rune]int {
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'0': 0,
	'A': -10,
	'<': 0,
	'V': 1,
	'>': 2,
	'^': 3,
}

var backwardsMap = map[int]rune {
	0: '<',
	1: 'v',
	2: '>',
	3: '^',
	-10: 'A',
}


type Point struct {
	row, col int
}

var keypad = [][]int {
	{7, 8, 9},
	{4, 5, 6},		//
	{1, 2, 3},		// 1, 2, 3
	{-1, 0, -10},	// x, 0, A
}

var rControl = [][]int { // TODO make this numeric too, so we can use same logic as for keypad
	{-1, 3, -10},	// x, ^, A -> use abs(-10) to be able to continue with same logic
	{0, 1, 2},		// <, V, >
}

func containsInRow(current, search int, slice [][]int) bool { // searches for the row where the current number is and looks if the search number is in the row
	for row, _ := range slice {
		for col, _ := range slice[row] {
			if slice[row][col] == current {
				for _, num := range slice[row] {
					if num == search {
						return true
					}
				}
			}
		}
	}
	return false
}

func getPos(num int, keypad [][]int) Point {
	for row := range keypad {
		for col := range keypad[row] {
			if keypad[row][col] == num {
				return Point{row, col}
			}
		}
	}
	return Point{-1, -1}
}

func horizontal(pos, pos2 int, keypad [][]int) []int {
	code := []int{}
	posCoord := getPos(pos, keypad)
	aCoord := getPos(pos2, keypad)
	for posCoord.col != aCoord.col {
		if posCoord.col > aCoord.col {
			code = append(code, 0) // move left
			posCoord.col--
		} else {
			code = append(code, 2) // move right
			posCoord.col++
		}
	}
	return code
}

func vertical(pos, pos2 int, keypad [][]int) []int {
	code := []int{}
	posCoord := getPos(pos, keypad)
	aCoord := getPos(pos2, keypad)
	for posCoord.row != aCoord.row {
		if posCoord.row > aCoord.row {
			code = append(code, 3) // move up
			posCoord.row--
		} else {
			code = append(code, 1) // move down
			posCoord.row++
		}
	}
	return code
}


func findInstructions(pos, pos2 int, keypad [][]int) []int {
    code := []int{}
    if containsInRow(pos, -1, keypad) {
		tmpCode := vertical(pos, pos2, keypad)
		code = append(code, tmpCode...)
		tmpCode = horizontal(pos, pos2, keypad)
		code = append(code, tmpCode...)
	} else {
		tmpCode := horizontal(pos, pos2, keypad)
		code = append(code, tmpCode...)
		tmpCode = vertical(pos, pos2, keypad)
		code = append(code, tmpCode...)
	}
	code = append(code, -10)
    return code
}

func typeKeypad(instructions []int, keypad [][]int, stringInstructions string) []int {
	pos := -10
	code := []int{}
	for _, instruction := range instructions {
		code = append(code, findInstructions(pos, instruction, keypad)...)
		pos = instruction
	}
	fmt.Print(stringInstructions, ": ")
	for _, char := range code {
		fmt.Print(string(backwardsMap[char]))
	}
	fmt.Println()
	// return all possible combinations + check for -1
	return code
}

func makeIntInstructions(code string) []int {
	intInstructions := make([]int, 0)
	for _, c := range code {
		intInstructions = append(intInstructions, keyMap[c])
	}
	return intInstructions
}

func main() {
    codes := []string{
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

	totalComplexity := 0

    for _, code := range codes {
        // fmt.Println("Processing code:", code)

        // Convert the code into instructions for the numeric keypad (Layer 1)
        intInstructions := makeIntInstructions(code)
        r1Instructions := typeKeypad(intInstructions, keypad, code)

        // Generate instructions for Robot 1 controlling the numeric keypad (Layer 2)
        r2Instructions := typeKeypad(r1Instructions, rControl, code)

        // Generate instructions for Robot 2 controlling Robot 1 (Layer 3)
        r3Instructions := typeKeypad(r2Instructions, rControl, code)

        // Extract numeric part
        numericPart, _ := strconv.Atoi(code[:len(code)-1]) // Remove 'A'

        // Calculate complexity for this code
        complexity := len(r3Instructions) * numericPart
        totalComplexity += complexity

        fmt.Printf("Code: %s, R3 Length: %d, Numeric: %d, Complexity: %d\n",
            code, len(r3Instructions), numericPart, complexity)
    }

    fmt.Println("Total Complexity:", totalComplexity)
}