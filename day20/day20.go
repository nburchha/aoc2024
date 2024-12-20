package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	SIZE = 71
	// FILE = "input/day20.txt"
	FILE = "input/testinput20"
)

type point struct{r, c int}

func bfs(grid []string, start, end point) int {
    visited := make(map[point]bool)
    queue := []struct {
        p point
        dist int
    }{{start, 0}}

    visited[start] = true
    dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        if curr.p == end {
            return curr.dist
        }

        for _, d := range dirs {
            nr, nc := curr.p.r+d.r, curr.p.c+d.c
            if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[nr]) && 
               rune(grid[nr][nc]) == '.' && !visited[point{nr, nc}] {
                visited[point{nr, nc}] = true
                queue = append(queue, struct { p point; dist int }{point{nr, nc}, curr.dist + 1})
            }
        }
    }
    return -1
}

func main() {
	input, _ := os.ReadFile(FILE)
	grid := strings.Split(string(input), "\n")
	var startPoint, endPoint point
	for row, line := range grid {
		for col, char := range line {
			if char == 'S' {
				grid[row] = grid[row][:col] + "." + grid[row][col+1:]
				startPoint = point{row, col}
			} else if char == 'E' {
				grid[row] = grid[row][:col] + "." + grid[row][col+1:]
				endPoint = point{row, col}
			}
		}
	}
	fmt.Println(startPoint, endPoint)
	fmt.Println(bfs(grid, startPoint, endPoint))
}