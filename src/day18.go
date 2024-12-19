package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

const (
	SIZE = 71
	FILE = "input/day18.txt"
	TESTFILE = "input/testinput18"
)

var seen = make(map[[2]int]bool)

func bfs(grid [][]rune, startRow, startCol int) int {
    type point struct{r, c int}
    visited := make(map[point]bool)
    queue := []struct {
        p point
        dist int
    }{{point{startRow, startCol}, 0}}

    visited[point{startRow, startCol}] = true
    dirs := []point{{1,0},{-1,0},{0,1},{0,-1}}

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        if curr.p.r == SIZE-1 && curr.p.c == SIZE-1 {
            return curr.dist
        }

        for _, d := range dirs {
            nr, nc := curr.p.r+d.r, curr.p.c+d.c
            if nr >= 0 && nr < SIZE && nc >= 0 && nc < SIZE && 
               grid[nr][nc] == '.' && !visited[point{nr,nc}] {
                visited[point{nr,nc}] = true
                queue = append(queue, struct{p point; dist int}{point{nr,nc}, curr.dist+1})
            }
        }
    }
    return -1
}

func calcSteps(grid [][]rune, row, col, count int) int {
    if row < 0 || row >= SIZE || col < 0 || col >= SIZE || grid[row][col] == '#' || seen[[2]int{row, col}] {
        return -1
    }

    if row == SIZE - 1 && col == SIZE - 1 {
        return count
    }

    seen[[2]int{row, col}] = true

    steps := -1
    directions := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
    for _, d := range directions {
        nr, nc := row + d[0], col + d[1]
        tmp := calcSteps(grid, nr, nc, count+1)
        if tmp != -1 && (steps == -1 || tmp < steps) {
            steps = tmp
        }
    }
    seen[[2]int{row, col}] = false
    return steps
}

func main() {
	input, _ := os.ReadFile(FILE)
	// input, _ := os.ReadFile(TESTFILE)
	lines := strings.Split(string(input), "\n")
	grid := make([][]rune, SIZE)
	for i := range grid {
		grid[i] = make([]rune, SIZE)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for index, line := range lines {
		re := regexp.MustCompile(`(\d+),(\d+)`)
		matches := re.FindAllStringSubmatch(line, -1)
		fmt.Println(matches)
		x, _ := strconv.Atoi(matches[0][1])
		y, _ := strconv.Atoi(matches[0][2])
		grid[x][y] = '#'
		
		if index >= 1024 {
			res := bfs(grid, 0, 0)
			if res == -1 {
				fmt.Println(line)
				fmt.Println(res)
				break
			}
		}
		
	}
	for line := range grid {
		fmt.Println(string(grid[line]))
	}
	fmt.Println(bfs(grid, 0, 0))
}