package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)

const (
	WIDE = 101
	TALL = 103
)

type robot struct {
	posX, posY, stepX, stepY int
}

func parseRobot(line string) (robot, error) {
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	if len(matches[0]) != 5 {
		return robot{}, fmt.Errorf("invalid line format")
	}
	// fmt.Println("matches:",matches)
	px, err := strconv.Atoi(matches[0][1])
	if err != nil {
		return robot{}, err
	}
	py, err := strconv.Atoi(matches[0][2])
	if err != nil {
		return robot{}, err
	}
	vx, err := strconv.Atoi(matches[0][3])
	if err != nil {
		return robot{}, err
	}
	vy, err := strconv.Atoi(matches[0][4])
	if err != nil {
		return robot{}, err
	}

	return robot{posX: px, posY: py, stepX: vx, stepY: vy}, nil
}

func parseInput(lines []string) []robot {
	result := make([]robot, 0)
	for _, line := range lines {
		rob, err := parseRobot(line)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			return nil
		}
		result = append(result, rob)
	}
	return result
}

func printGrid(robots []robot) {
	grid := make([][]rune, TALL)
	for i := 0; i < TALL; i++ {
		grid[i] = make([]rune, WIDE)
		for j := 0; j < WIDE; j++ {
			grid[i][j] = ' '
		}
	}
	for _, rob := range robots {
		// if grid[rob.posY][rob.posX] != '.' {
		// 	grid[rob.posY][rob.posX] += 1
		// } else {
			grid[rob.posY][rob.posX] = '1'
		// }
	}
	for i := 0; i < TALL; i++ {
		// if i == TALL / 2 {
		// 	fmt.Println()
		// } else {
			fmt.Println(string(grid[i]))
		// }
	}
}

func mod(a, b int) int {
    return (a%b + b) % b
}

func getQuadrant(rob robot) int {
	if rob.posX < WIDE/2 && rob.posY < TALL/2 {
		return 0
	} else if rob.posX > WIDE/2 && rob.posY < TALL/2 {
		return 1
	} else if rob.posX < WIDE/2 && rob.posY > TALL/2 {
		return 2
	} else if rob.posX > WIDE/2 && rob.posY > TALL/2 {
		return 3
	}
	return -1
}

func multiplyEachQuadrant(robots []robot) int {
	quadrants := make(map[int]int)
	for _, rob := range robots {
		// fmt.Println("Robot:", rob, "Quadrant:", getQuadrant(rob))
		switch getQuadrant(rob) {
		case 0:
			quadrants[0] += 1
		case 1:
			quadrants[1] += 1
		case 2:
			quadrants[2] += 1
		case 3:
			quadrants[3] += 1
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

// tree is there when the robots are not overlapping
func probableChristmasTree(robots []robot) bool {
	pos := make(map[[2]int]int)
	for _, rob := range robots {
		pos[[2]int{rob.posX, rob.posY}] += 1
	}
	for _, count := range pos {
		if count > 1 {
			return false
		}
	}
	return true
}

func main() {
	input, _ := os.ReadFile("input/day14.txt")
	// input, _ := os.ReadFile("input/testinput14")
	blocks := strings.Split(string(input), "\n\n")
	lines := strings.Split(blocks[0], "\n")
	robots := parseInput(lines)

	for i:=0; i<10000; i++ {
		if probableChristmasTree(robots) {
			fmt.Println("PART2 - Second:", i)
		}
		for index := range robots {
			robots[index].posX = mod(robots[index].posX + robots[index].stepX, WIDE)
			robots[index].posY = mod(robots[index].posY + robots[index].stepY, TALL)
		}
	}
	fmt.Println("PART1:",multiplyEachQuadrant(robots))
}