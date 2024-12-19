package main

import (
	"fmt"
	"os"
	"strings"
)

type Side struct {
	boundary1, boundary2 [2]int
	dir int
}

var seen = make(map[[2]int]byte)
var directions = [][]int{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

func checkDir(grid []string, row, col, dir int, plantType byte) bool {
	switch dir {
	case 0:
		row += 1
	case 1:
		col += 1
	case 2:
		row -= 1
	case 3:
		col -= 1
	}

	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) || grid[row][col] != plantType {
		return true
	}
	return false
}

func getSide(grid []string, row, col, dir int, plantType byte) (Side, error) {
	rows, cols := len(grid), len(grid[0])
	startRow, startCol := row, col
	endRow, endCol := row, col

	switch dir {
	case 0:
		for startCol > 0 && grid[row][startCol-1] == plantType && checkDir(grid, row, startCol-1, dir, plantType) { startCol-- }
		for endCol < cols-1 && grid[row][endCol+1] == plantType && checkDir(grid, row, endCol+1, dir, plantType) { endCol++ }
	case 1:
		for startRow > 0 && grid[startRow-1][col] == plantType && checkDir(grid, startRow-1, col, dir, plantType) { startRow-- }
		for endRow < rows-1 && grid[endRow+1][col] == plantType && checkDir(grid, endRow+1, col, dir, plantType) { endRow++ }
	case 2:
		for startCol > 0 && grid[row][startCol-1] == plantType && checkDir(grid, row, startCol-1, dir, plantType) { startCol-- }
		for endCol < cols-1 && grid[row][endCol+1] == plantType && checkDir(grid, row, endCol+1, dir, plantType) { endCol++ }
	case 3:
		for startRow > 0 && grid[startRow-1][col] == plantType && checkDir(grid, startRow-1, col, dir, plantType) { startRow-- }
		for endRow < rows-1 && grid[endRow+1][col] == plantType && checkDir(grid, endRow+1, col, dir, plantType) { endRow++ }
	}

	if checkDir(grid, row, col, dir, plantType) {
		return Side{
			boundary1: [2]int{startRow, startCol},
			boundary2: [2]int{endRow, endCol},
			dir: dir,
		}, nil
	}

	return Side{}, fmt.Errorf("no valid side found")
}

func sideExists(side Side, sides []Side) bool {
	for _, s := range sides {
		if s.dir == side.dir {
			switch s.dir%2 {
			case 0: // gerade zahl -> 0 2 -> top down
				if side.boundary1[0] == s.boundary1[0] && side.boundary1[1] >= s.boundary1[1] && side.boundary2[1] <= s.boundary2[1] { return true }
			case 1:
				if side.boundary1[1] == s.boundary1[1] && side.boundary1[0] >= s.boundary1[0] && side.boundary2[0] <= s.boundary2[0] { return true }
			}
		}
	}
	return false
}

func checkSurroundings2(grid []string, row, col int, sides []Side) []Side {
	plantType := grid[row][col]
	for i:=0;i<4;i++ {
		side, err := getSide(grid, row, col, i, plantType)
		if err != nil {
			continue
		}
		if !sideExists(side, sides) {
			sides = append(sides, side)
		}
	}
	return sides
}


func getCostOfArea2(grid []string, row, col int, plantType byte, sides []Side) (int, []Side) {
	if sides == nil {
		sides = []Side{}
	}
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) || grid[row][col] != plantType {
		return 0, sides
	} else if _, err := seen[[2]int{row,col}]; err { // if this element exists in the map
		return 0, sides
	}
	seen[[2]int{row, col}] = plantType
	var count int
	for _, dir := range directions {

		tmpCount, tmpSides := getCostOfArea2(grid, row+dir[0], col+dir[1], plantType, sides)
		count += tmpCount
		sides = tmpSides
	}

	sides = checkSurroundings2(grid, row, col, sides)
	return count+1, sides
}

func iterateGrid2(grid []string) int {
	var totalCount int
	for row := range grid {
		for col := range grid[row] {
			count, sides := getCostOfArea2(grid, row, col, grid[row][col], nil)
			totalCount += len(sides) * count
		}
	}
	return totalCount
}

func main() {
	input, _ := os.ReadFile("input/day12.txt")
	// input, _ := os.ReadFile("input/testinput12")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	lines := strings.Split(split[0], "\n")
	fmt.Println(iterateGrid2(lines))
}