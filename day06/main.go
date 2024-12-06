package main

import (
	"fmt"
	"os"
	"bufio"
)

type Coordinate struct {
	X, Y, Dir int
}

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
	return lines, nil
}

func findStart(input []string) Coordinate {
	for row, line := range input {
		for col, char := range line {
			if char == '^' {
				return Coordinate{X: col, Y: row, Dir: 0}
			}
		}
	}
	return Coordinate{X: -1, Y: -1, Dir: -1}
}

func getNextPos(start Coordinate) Coordinate {
	next := start
	switch start.Dir {
	case 0:
		next.Y--
	case 1:
		next.X++
	case 2:
		next.Y++
	case 3:
		next.X--
	}
	return next
}

func outOfBounds(next Coordinate, input []string) bool {
	return next.X < 0 || next.Y < 0 || next.X >= len(input[0]) || next.Y >= len(input)
}

func calcDistinctCoordinates(input []string, start Coordinate) {
	for {
		if outOfBounds(start, input) {
			break
		}
		next := getNextPos(start)
		if !outOfBounds(next, input) && input[next.Y][next.X] == '#' {
			start.Dir = (start.Dir + 1) % 4
		}
		input[start.Y] = input[start.Y][:start.X] + "X" + input[start.Y][start.X+1:]
		start = getNextPos(start)
	}
}

func countX(input []string) int {
	count := 0
	for _, line := range input {
		for _, char := range line {
			if char == 'X' {
				count++
			}
		}
	}
	return count
}

func makeCopy(input []string) []string {
	tmp := make([]string, len(input))
	copy(tmp, input)
	return tmp
}

func tryEndlessLoop(input []string, start Coordinate) bool {
	flag := false
	visited:= make(map[Coordinate]bool)
	for {
		if outOfBounds(start, input) {
			break
		}
		if visited[start] {
			flag = true
			break
		}
		visited[start] = true
		next := getNextPos(start)
		if !outOfBounds(next, input) && (input[next.Y][next.X] == '#' || input[next.Y][next.X] == 'O') {
			start.Dir = (start.Dir + 1) % 4
		} else {
			start = next
		}
	}
	return flag
}

func part2(input []string, start Coordinate) int {
	count := 0
	for row, line := range input {
		for col := range line {
			tmp := makeCopy(input)
			tmp[row] = tmp[row][:col] + "O" + input[row][col+1:]
			if tryEndlessLoop(tmp, start) {
				count++
			}
		}
	}
	return count
}

func main() {
	input, err := parseInput("input/day06.txt")
	if (err != nil) {
		fmt.Println(err)
		return
	}
	start := findStart(input)
	if start.X == -1 || start.Y == -1 {
		fmt.Println("start not found")
	}
	input1 := makeCopy(input)
	calcDistinctCoordinates(input1, start)
	fmt.Println("part1:", countX(input1))
	fmt.Println("part2:", part2(input, start))
}