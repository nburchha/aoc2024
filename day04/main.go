package main

import (
	"os"
	"fmt"
	"bufio"
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

/*
search for all XMAS in input
they can be left to right, right to left, bottom to top, top bottom, diagonal
can overlap

iterate over each character of each line and always give whole slice of strings, from each char start recursion loop which searchs
*/

// direction: 0 left right; 1 right left; 2 up down; 3 down up; 4 left diag up; 5 left diag down; 6 right diag up; 7 right diag down
func recSearch(arr []string, x int, y int, direction int, searchIndex int) bool {
	if y < 0 || y >= len(arr) || x < 0 || x >= len(arr[y]) || arr[y][x] != "XMAS"[searchIndex] {
		return false
	}
	if searchIndex == 3 && arr[y][x] == "XMAS"[searchIndex] {
		return true
	}
	switch direction {
	case 0:
		return recSearch(arr, x + 1, y, direction, searchIndex + 1)
	case 1:
		return recSearch(arr, x - 1, y, direction, searchIndex + 1)
	case 2:
		return recSearch(arr, x, y - 1, direction, searchIndex + 1)
	case 3:
		return recSearch(arr, x, y + 1, direction, searchIndex + 1)
	case 4:
		return recSearch(arr, x + 1, y + 1, direction, searchIndex + 1)
	case 5:
		return recSearch(arr, x + 1, y - 1, direction, searchIndex + 1)
	case 6:
		return recSearch(arr, x - 1, y + 1, direction, searchIndex + 1)
	case 7:
		return recSearch(arr, x - 1, y - 1, direction, searchIndex + 1)
	default:
		return false
	}
}

func searchX(arr []string, x int, y int) bool {
	if y < 1 || y >= len(arr)-1 || x < 1 || x >= len(arr[y])-1 || arr[y][x] != 'A' {
		return false
	}
	slices := []string{"MSMS", "MMSS", "SMSM", "SSMM"}
	for _, slice := range slices {
		if arr[y+1][x-1] == slice[0] && arr[y+1][x+1] == slice[1] && arr[y-1][x-1] == slice[2] && arr[y-1][x+1] == slice[3] {
			return true
		}
	}
	return false
}

func searchXMAS(arr []string, x int, y int) int {
	count := 0
	for dir:=0;dir<8;dir++ {
		if recSearch(arr, x, y, dir, 0) {
			count++
		}
	}
	return count
}

func main() {
	lines, err := parseInput("input/day04.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	count := 0
	count2 := 0
	for y :=0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			count += searchXMAS(lines, x, y)
			if searchX(lines, x, y) {
				count2++
			}
		}
	}
	fmt.Println("part1:",count)
	fmt.Println("part2:",count2)
}