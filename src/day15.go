package main

import (
	"fmt"
	"os"

	// "strconv"
	"strings"
	// "regexp"
)

type Coord struct {
	x, y int
}

func parseMapAndMoves(lines []string) ([]string, []Coord, Coord) {
	mapData := make([]string, 0)
	robotMoves := make([]Coord, 0)
	rPos := Coord{x: -1, y: -1}
	parseMoves := false
	for row, line := range lines { // parse map
		fmt.Println(line)
		if line == "" {
			parseMoves = true
			continue
		}
		if parseMoves {
			for _, c := range line {
				switch c {
				case '^':
					robotMoves = append(robotMoves, Coord{x: 0, y: -1})
				case 'v':
					robotMoves = append(robotMoves, Coord{x: 0, y: 1})
				case '>':
					robotMoves = append(robotMoves, Coord{x: 1, y: 0})
				case '<':
					robotMoves = append(robotMoves, Coord{x: -1, y: 0})
				}
			}
		} else {
			for col, c := range line {
				if c == '@' {
					rPos = Coord{x: col, y: row}
				}
			}

			mapData = append(mapData, line)
		}
	}
	return mapData, robotMoves, rPos
}

func replaceCharInMap(mapData []string, x, y int, newChar rune) {
	mapRow := []rune(mapData[y])
	mapRow[x] = newChar
	mapData[y] = string(mapRow)
}

func move(c byte, pos Coord, dir Coord, mapData []string) (bool, Coord) {
	flag := false
	if pos.y + dir.y < 0 || pos.y + dir.y >= len(mapData) || pos.x + dir.x < 0 || pos.x + dir.x >= len(mapData[pos.y + dir.y]) || mapData[pos.y + dir.y][pos.x + dir.x] == '#' {
		return false, pos
	} else if mapData[pos.y + dir.y][pos.x + dir.x] == 'O' {
		flag, _ = move(mapData[pos.y + dir.y][pos.x + dir.x], Coord{x: pos.x + dir.x, y: pos.y + dir.y}, dir, mapData)
	} else {
		flag = true
	}
	if flag {
		replaceCharInMap(mapData, pos.x, pos.y, '.')
		replaceCharInMap(mapData, pos.x + dir.x, pos.y + dir.y, rune(c))
		pos.x += dir.x
		pos.y += dir.y
	}
	return flag, pos
}

func printMap(mapData []string) {
	for _, line := range mapData {
		fmt.Println(line)
	}
}

func calcSumOfBoxes(mapData []string) int {
	sum := 0
	for row, line := range mapData {
		for col, c := range line {
			if c == 'O' {
				sum += (row)*100 + (col)
			}
		}
	}
	return sum
}

func main() {
	input, _ := os.ReadFile("input/day15.txt")
	// input, _ := os.ReadFile("input/testinput15")
	lines := strings.Split(string(input), "\n")
	mapData, robotMoves, rPos := parseMapAndMoves(lines)
	fmt.Println(mapData, robotMoves, rPos)

	for _, dir := range robotMoves {
		_, rPos = move('@', rPos, dir, mapData)
		// fmt.Println("count:", count, dir)
		// printMap(mapData)
	}

	fmt.Println(calcSumOfBoxes(mapData))
}