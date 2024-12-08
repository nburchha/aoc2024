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

func checkAntiNode(antinode, b Coord, rows, cols int) bool {
	return antinode.X >= 0 && antinode.X < cols && antinode.Y >= 0 && antinode.Y < rows && antinode != b
}

func validAntiNode(a, b Coord, rows, cols int) (int, []Coord) {
	var delta Coord
	delta.X = b.X - a.X
	delta.Y = b.Y - a.Y
	// fmt.Println("coords:", a, b)
	antinode1 := Coord{X: b.X + delta.X, Y: b.Y + delta.Y}
	antinode2 := Coord{X: a.X - delta.X, Y: a.Y - delta.Y}

	var count int
	res := []Coord{
		{X: -1, Y: -1},
		{X: -1, Y: -1},
	}
	if checkAntiNode(antinode1, a, rows, cols) {
		fmt.Println("valid antinode:", antinode1)
		count++
		res[0] = antinode1
	}
	if checkAntiNode(antinode2, b, rows, cols) {
		fmt.Println("valid antinode:", antinode2)
		count++
		res[1] = antinode2
	}
	// fmt.Println("count",count)
	return count, res
}

func inBoundAntiNodes(antennas []Coord, rows, cols int) (int, []Coord) {
	var count int
	var antiNodes []Coord
	for i, a := range antennas {
		for j := i + 1; j < len(antennas); j++ {
			b := antennas[j]
			if a == b { continue }
			tmp, positions := validAntiNode(a, b, rows, cols)
			count += tmp
			if count != 0 {
				antiNodes = append(antiNodes, positions...)
			}
		}
	}
	// for _, line := range grid {
	// 	fmt.Println(line)
	// }
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
	fmt.Println(len(antiNodes) - 1)
}
