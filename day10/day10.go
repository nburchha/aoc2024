package main

import (
	"os"
	"strings"
	"fmt"
	"strconv"
)

func convertToInt(input []string) [][]int {
	var result [][]int
	for _, str := range input {
		var row []int
		for _, char := range str {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil
			}
			row = append(row, num)
		}
		result = append(result, row)
	}
	return result
}

func recSearchScore1(trailMap [][]int, x, y, score int, count *int, visited map[[2]int]bool) {
	if y < 0 || y >= len(trailMap) || x < 0 || x >= len(trailMap[y]) || trailMap[y][x] != score+1 || visited[[2]int{y,x}] {
		return
	}
	visited[[2]int{y, x}] = true
	if trailMap[y][x] == 9 {
		*count += 1
		return
	}
	recSearchScore1(trailMap, x+1, y, score+1, count, visited)
	recSearchScore1(trailMap, x-1, y, score+1, count, visited)
	recSearchScore1(trailMap, x, y+1, score+1, count, visited)
	recSearchScore1(trailMap, x, y-1, score+1, count, visited)
}

func recSearchScore2(trailMap [][]int, x, y, score int, count *int) {
	if y < 0 || y >= len(trailMap) || x < 0 || x >= len(trailMap[y]) || trailMap[y][x] != score+1 {
		return
	}
	if trailMap[y][x] == 9 {
		*count += 1
		return
	}
	recSearchScore2(trailMap, x+1, y, score+1, count)
	recSearchScore2(trailMap, x-1, y, score+1, count)
	recSearchScore2(trailMap, x, y+1, score+1, count)
	recSearchScore2(trailMap, x, y-1, score+1, count)
}

func main() {
	input, _ := os.ReadFile("input/day10.txt")
	// input, _ := os.ReadFile("input/testinput10")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	lines := strings.Split(split[0], "\n")
	trailMap := convertToInt(lines)

	var count1 int
	var count2 int
	for row := range trailMap {
		for col := range trailMap[row] {
			if trailMap[row][col] == 0 {
				visited := make(map[[2]int]bool)
				recSearchScore1(trailMap, col, row, -1, &count1, visited)
				recSearchScore2(trailMap, col, row, -1, &count2)
			}
		}
	}
	fmt.Println("part1:",count1,"part2:",count2)
}