package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	X, Y int
}

func ParseInput(path string) ([]Pair, [][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pairs []Pair

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		pairs = append(pairs, Pair{X: x, Y: y})
	}

	var matrix [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		update := make([]int, len(parts))
		for i, num := range parts {
			update[i], _ = strconv.Atoi(num)
		}
		matrix = append(matrix, update)
	}

	return pairs, matrix, nil
}

func validLine(line []int, pairs []Pair) bool {
	index := make(map[int]int)
	for i, num := range line {
		index[num] = i
	}
	for _, pair := range pairs {
		if idxX, okX := index[pair.X]; okX {
			if idxY, okY := index[pair.Y]; okY && idxX > idxY {
				return false
			}
		}
	}
	return true
}

func reorderLine(line []int, pairs []Pair) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, num := range line {
		graph[num] = []int{}
		inDegree[num] = 0
	}
	for _, pair := range pairs {
		if contains(line, pair.X) && contains(line, pair.Y) {
			graph[pair.X] = append(graph[pair.X], pair.Y)
			inDegree[pair.Y]++
		}
	}

	var sorted []int
	queue := []int{}
	for num, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, num)
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return sorted
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func iterateMatrix(matrix [][]int, pairs []Pair) (int, int) {
	sum1 := 0
	sum2 := 0
	for _, line := range matrix {
		if !validLine(line, pairs) {
			line = reorderLine(line, pairs)
			mid := len(line) / 2
			sum2 += line[mid]
		} else {
			mid := len(line) / 2
			sum1 += line[mid]
		}
	}
	return sum1, sum2
}

func main() {
	pairs, matrix, err := ParseInput("input/day05.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	part1, part2 := iterateMatrix(matrix, pairs)
	fmt.Println("Sum of corrected middle values (part2):", part1)
	fmt.Println("Sum of corrected middle values (part2):", part2)
}
