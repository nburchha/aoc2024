package main

import (
	"fmt"
	"os"
	"strings"
)

var seen = make(map[[2]int]byte)
var directions = [][]int{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

func checkSurroundings(grid []string, row, col int) int {
	count := 0
	plantType := grid[row][col]
	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]
		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[newRow]) {
			if grid[newRow][newCol] != plantType {
				count++
			}
		} else {
			count++
		}
	}
	return count
}


func getCostOfArea(grid []string, row, col int, plantType byte) (int, int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) || grid[row][col] != plantType {
		return 0, 0
	} else if _, err := seen[[2]int{row,col}]; err { // if this element exists in the map
		return 0, 0
	}
	seen[[2]int{row, col}] = plantType
	var count, cost int
	for _, dir := range directions {
		tmpCost, tmpCount := getCostOfArea(grid, row+dir[0], col+dir[1], plantType)
		count += tmpCount
		cost += tmpCost
	}

	cost += checkSurroundings(grid, row, col)

	return cost, count+1
}

func iterateGrid(grid []string) int {
	var totalCount int
	for row := range grid {
		for col := range grid[row] {
			cost, count := getCostOfArea(grid, row, col, grid[row][col])
			totalCount += cost * count
			// fmt.Println((string)(grid[row][col]),":",cost,count)
		}
	}
	return totalCount
}

func main() {
	input, _ := os.ReadFile("input/day12.txt")
	// input, _ := os.ReadFile("input/testinput12")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	lines := strings.Split(split[0], "\n")
	fmt.Println(iterateGrid(lines))
}