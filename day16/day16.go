package main

import (
	"fmt"
	"os"

	// "strconv"
	"strings"
	// "regexp"
)

const (
	PART2 = false
)

type Coord struct {
	x, y int
}

var visited = make(map[[2]Coord]int)


func calcScore(dir, prevDir Coord) int {
	if dir.x == prevDir.x && dir.y == prevDir.y {
		return 1
	} else if dir.x == -prevDir.x || dir.y == -prevDir.y { // 180 degree turn
		return 2001
	}
	return 1001
}

func searchBestMoves(grid []string, current, end, dir Coord, score int) (int, map[Coord]bool) {
	if current == end {
		return score, map[Coord]bool{current: true}
	} else if current.x < 0 || current.y < 0 || current.y >= len(grid) || current.x >= len(grid[current.y]) || grid[current.y][current.x] == '#' {
		return -1, nil
	}
	if visitedScore, ok := visited[[2]Coord{current, dir}]; ok {
		if visitedScore < score {
			return -1, nil
		}
	}
	visited[[2]Coord{current, dir}] = score
	var dirs = []Coord{{x: 0, y: -1}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 1, y: 0}}
	best := -1
	bestTiles := map[Coord]bool{}
	for _, d := range dirs {
		tmpScore, tmpTiles := searchBestMoves(grid, Coord{x: current.x + d.x, y: current.y + d.y}, end, d, score+calcScore(dir, d))
		if tmpScore != -1 && (best == -1 || tmpScore < best) {
			best = tmpScore
			bestTiles = tmpTiles
		} else if tmpScore != -1 && tmpScore == best {
			for tile := range tmpTiles {
				bestTiles[tile] = true
			}
		}
	}
	if best == -1 {
		return -1, nil
	}
	bestTiles[current] = true
	return best, bestTiles
}

func main() {
	input, _ := os.ReadFile("input/day16.txt")
	// input, _ := os.ReadFile("input/testinput16")
	grid := strings.Split(string(input), "\n")
	var start, end Coord
	for row, line := range grid {
		for col, c := range line {
			if c == 'E' {
				end = Coord{x: col, y: row}
			} else if c == 'S' {
				start = Coord{x: col, y: row}
			}
		}
	}
	score, tileMap := searchBestMoves(grid, start, end, Coord{x: 1, y: 0}, 0)
	fmt.Println(score, len(tileMap))
	// for row, line := range grid {
	// 	for col := range line {
	// 		if _, ok := tileMap[Coord{x: col, y: row}]; ok {
	// 			fmt.Print("O")
	// 		} else {
	// 			fmt.Print(string(grid[row][col]))
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}