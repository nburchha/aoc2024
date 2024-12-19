package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

const (
	COST_A = int64(3)
	COST_B = int64(1)
	MAXDEPTH = int64(100)
)

type stuff struct {
	aX, aY, bX, bY, pX, pY int64
}


func parseLine(line, searchX, searchY string) (int64, int64) {
	parts := strings.Split(line, " ")
	searchIndex := 2
	if len(parts) == 4 {
		searchIndex = 3
	}
	xStr := strings.TrimPrefix(parts[searchIndex-1], searchX)
	xStr = strings.TrimSuffix(xStr, ",")
	x, _ := strconv.ParseInt(xStr, 10, 64)

	yStr := strings.TrimPrefix(parts[searchIndex], searchY)
	y, _ := strconv.ParseInt(yStr, 10, 64)

	return x, y
}

func parseInput(blocks []string) []stuff {
	result := make([]stuff, len(blocks) - 1)
	fmt.Println(len(blocks) -1)
	for i, block := range blocks {
		if i == len(blocks) - 1 { break }
		lines := strings.Split(block, "\n")
		aX, aY := parseLine(lines[0], "X+", "Y+")
		bX, bY := parseLine(lines[1], "X+", "Y+")
		pX, pY := parseLine(lines[2], "X=", "Y=")
		// pX += 10000000000000
		// pY += 10000000000000
		result[i] = stuff{aX, aY, bX, bY, pX, pY}
	}
	return result
}

// part1
func calcCheapestPrize(data stuff) int {
	var cost int
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if int(data.aX) * i + int(data.bX) * j == int(data.pX) && int(data.aY) * i + int(data.bY) * j == int(data.pY) {
				if i*3 + j < cost || cost == 0 {
					cost = i*int(COST_A) + j*int(COST_B)
				}
			}
		}
	}
	return cost
}

func play(g stuff, maxpress int64) (int64, bool) {
	ax, ay := float64(g.aX), float64(g.aY)
	bx, by := float64(g.bX), float64(g.bY)
	px, py := float64(g.pX), float64(g.pY)

	j := (py*bx/by - px) / (ay*bx/by - ax)
	k := (px - ax*j) / bx

	if j < 0 || k < 0 {
		return -1, false
	}

	ji, ki := int64(math.Round(j)), int64(math.Round(k))

	iswin := ji*g.aX + ki*g.bX == g.pX &&
		ji*g.aY + ki*g.bY == g.pY &&
		ji <= maxpress &&
		ki <= maxpress
	return ji*COST_A + ki*COST_B, iswin
}


func main() {
	input, _ := os.ReadFile("input/day13.txt")
	// input, _ := os.ReadFile("input/testinput13")
	blocks := strings.Split(string(input), "\n\n")
	parsedData := parseInput(blocks)

	// Part 1
	var part1 int
	var part2 int64
	var count int
	for _, data := range parsedData {
		// fmt.Println("data:",data)
		part1 += calcCheapestPrize(data)
		data.pX += 10000000000000
		data.pY += 10000000000000
		tmp2, ok := play(data, 99999999999999999)
		if ok {
			part2 += tmp2
			count++
		}
	}
	fmt.Println("Total count part1:", part1)
	fmt.Println("Total count part2:", part2)
}