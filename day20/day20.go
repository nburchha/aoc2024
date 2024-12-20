package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	SIZE = 71
	FILE = "input/day20.txt"
	// FILE = "input/testinput20"
)

type point struct{r, c int}

var dirs = []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func isTrack(ch byte) bool {
    return ch == '.' || ch == 'S' || ch == 'E'
}

// allDistances performs a BFS from a single start point and returns a 2D slice of distances to all reachable cells.
// If a cell is not reachable, distance will be -1.
func allDistances(grid []string, start point) [][]int {
    h := len(grid)
    w := len(grid[0])
    dist := make([][]int, h)
    for i := range dist {
        dist[i] = make([]int, w)
        for j := range dist[i] {
            dist[i][j] = -1
        }
    }

    dist[start.r][start.c] = 0
    q := []point{start}
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        for _, d := range dirs {
            nr, nc := cur.r+d.r, cur.c+d.c
            if nr < 0 || nr >= h || nc < 0 || nc >= w {
                continue
            }
            if dist[nr][nc] == -1 && isTrack(grid[nr][nc]) {
                dist[nr][nc] = dist[cur.r][cur.c] + 1
                q = append(q, point{nr,nc})
            }
        }
    }
    return dist
}

// solvePart2 implements the logic for part two.
func solvePart2(grid []string, start, end point) int {
    h := len(grid)
    w := len(grid[0])

    normalDistFromStart := allDistances(grid, start)
    normalDistFromEnd := allDistances(grid, end)

    normalTime := normalDistFromStart[end.r][end.c]

    if normalTime == -1 {
        return 0
    }

    type cheatID struct {
        sr, sc, er, ec int
    }
    bestSavings := make(map[cheatID]int)

    hIn := func(r, c int) bool {return r>=0 && r<h && c>=0 && c<w}

    for sr := 0; sr < h; sr++ {
        for sc := 0; sc < w; sc++ {
            if normalDistFromStart[sr][sc] == -1 {
                continue
            }
            distStart := normalDistFromStart[sr][sc]

            visited := make([][][]bool, h)
            for i := 0; i < h; i++ {
                visited[i] = make([][]bool, w)
                for j := 0; j < w; j++ {
                    visited[i][j] = make([]bool, 21)
                }
            }

            // BFS queue for cheat mode
            type state struct {
                r, c, steps int
            }

            visited[sr][sc][0] = true
            queue := []state{{sr, sc, 0}}

            for len(queue) > 0 {
                cur := queue[0]
                queue = queue[1:]

                for _, d := range dirs {
                    nr, nc := cur.r+d.r, cur.c+d.c
                    if !hIn(nr,nc) {
                        continue
                    }
                    nsteps := cur.steps + 1
                    if nsteps > 20 {
                        continue
                    }
                    if !visited[nr][nc][nsteps] {
                        visited[nr][nc][nsteps] = true
                        queue = append(queue, state{nr,nc,nsteps})
                    }

                    // Potential cheat end conditions:
                    // If we reached 'E':
                    if grid[nr][nc] == 'E' {
                        cheatedTime := distStart + nsteps
                        saving := normalTime - cheatedTime
                        cid := cheatID{sr, sc, nr, nc}
                        if saving > bestSavings[cid] {
                            bestSavings[cid] = saving
                        }
                    } else {
                        ch := grid[nr][nc]
                        if isTrack(ch) && normalDistFromEnd[nr][nc] != -1 {
                            remainder := normalDistFromEnd[nr][nc]
                            cheatedTime := distStart + nsteps + remainder
                            saving := normalTime - cheatedTime
                            cid := cheatID{sr, sc, nr, nc}
                            if saving > bestSavings[cid] {
                                bestSavings[cid] = saving
                            }
                        }
                    }
                }
            }
        }
    }

    count := 0
    for _, s := range bestSavings {
        if s >= 100 {
            count++
        }
    }
    return count
}

func bfs(grid []string, start, end point) int {
    visited := make(map[point]bool)
    queue := []struct {
        p point
        dist int
    }{{start, 0}}

    visited[start] = true

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
	var count int
	originalTime := bfs(grid, startPoint, endPoint)
	fmt.Println("original:", originalTime)
	for row, line := range grid {
		for col, char := range line {
			if char == '#' {
				grid[row] = grid[row][:col] + "." + grid[row][col+1:]
				time := bfs(grid, startPoint, endPoint)
				if originalTime - time >= 100 {
					count++
				}
				grid[row] = grid[row][:col] + "#" + grid[row][col+1:]
			}
		}
	}
	fmt.Println("Part1:", count)
    fmt.Println("Part 2:", solvePart2(grid, startPoint, endPoint))
}