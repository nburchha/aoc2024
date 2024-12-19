package main

import (
	"fmt"
	// "bufio"
	"os"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func checkAntiNode1(antinode, b Coord, rows, cols int) bool {
	return antinode.X >= 0 && antinode.X < cols && antinode.Y >= 0 && antinode.Y < rows && antinode != b
}

func checkAntiNode2(antinode, b Coord, rows, cols int) bool {
	return antinode.X >= 0 && antinode.X < cols && antinode.Y >= 0 && antinode.Y < rows
}

func validAntiNode(a, b Coord, rows, cols int) (int, []Coord) {
	var delta Coord
	delta.X = b.X - a.X
	delta.Y = b.Y - a.Y

	count := 0
	res := make([]Coord, 0)
	res = append(res, a)
	res = append(res, b)
	for i := 1; i < 1000; i++ {
		antinode := Coord{X: (b.X + delta.X*i), Y: (b.Y + delta.Y*i)}
		if checkAntiNode2(antinode, a, rows, cols) {
			fmt.Println("valid antinode:", antinode)
			count++
			res = append(res, antinode)
		}
		antinode = Coord{X: (b.X - delta.X*i), Y: (b.Y - delta.Y*i)}
		if checkAntiNode2(antinode, a, rows, cols) {
			fmt.Println("valid antinode:", antinode)
			count++
			res = append(res, antinode)
		}
		antinode = Coord{X: (a.X + delta.X*i), Y: (a.Y + delta.Y*i)}
		if checkAntiNode2(antinode, b, rows, cols) {
			fmt.Println("valid antinode:", antinode)
			count++
			res = append(res, antinode)
		}
	}
	return count, res
}

func inBoundAntiNodes(antennas []Coord, rows, cols int) (int, []Coord) {
	var count int
	var antiNodes []Coord
	for i, a := range antennas {
		for j := i + 1; j < len(antennas); j++ {
			b := antennas[j]
			if a == b {
				continue
			}
			tmp, positions := validAntiNode(a, b, rows, cols)
			count += tmp
			if count != 0 {
				antiNodes = append(antiNodes, positions...)
			}
		}
	}
	return count, antiNodes
}

func removeDuplicates(coords []Coord) []Coord {
	uniqueMap := make(map[Coord]struct{})
	for _, coord := range coords {
		uniqueMap[coord] = struct{}{}
	}

	uniqueCoords := make([]Coord, 0, len(uniqueMap))
	for coord := range uniqueMap {
		uniqueCoords = append(uniqueCoords, coord)
	}

	return uniqueCoords
}

func main() {
	// input, _ := os.ReadFile("input/testinput8")
	input, _ := os.ReadFile("input/day08.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	fmt.Println(split[0])
	rows := len(strings.Split(split[0], "\n"))
	cols := len(strings.Split(split[0], "\n")[0])
	antennas := make(map[rune][]Coord)
	for row, line := range strings.Split(split[0], "\n") {
		for col := range line {
			if line[col] != '.' {
				antennas[rune(line[col])] = append(antennas[rune(line[col])], Coord{col, row})
			}
		}
	}
	var count int
	var antiNodes []Coord
	for _, v := range antennas {
		tmp, tmp1 := inBoundAntiNodes(v, rows, cols)
		count += tmp
		antiNodes = append(antiNodes, tmp1...)
	}
	antiNodes = removeDuplicates(antiNodes)
	fmt.Println(len(antiNodes))
}
