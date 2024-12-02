package main

import (
	"fmt"
	"aoc2024/utils"
)

// rules:
// - either all decreasing or all increasing
// - any two adjacent levels differ by at least one and at most three -> no bigger steps than 3, 
func calcSafeOrUnsafe(report []int) (bool){
	priorLevel := report[0]
	increasing, decreasing := 1, 1
	safeDistance := true
	for i:=1; i<len(report); i++ {
		fmt.Println("i:", i, "priorLevel:", priorLevel, "report[i]:", report[i])
		if report[i] < priorLevel {
			decreasing += 1
		} else if report[i] > priorLevel {
			increasing += 1
		}
		if priorLevel - report[i] > 3 || report[i] - priorLevel > 3 || priorLevel == report[i] {
			safeDistance = false
		}
		priorLevel = report[i]
	}
	if (increasing == len(report) || decreasing == len(report)) && safeDistance{
		fmt.Println("Safe, increasing:", increasing, "decreasing:", decreasing, "length:", len(report))
		return true
	} else {
		fmt.Println("Unsafe, increasing:", increasing, "decreasing:", decreasing, "length:", len(report))
		return false
	}
}

func loopInput(input [][]int) {
	count := 0
	for _, report := range input {
		fmt.Println(report)
		if calcSafeOrUnsafe(report) {
			count += 1
		}
	}
	fmt.Println("Safe reports:", count)
}

func main() {
	input, err := utils.ParseInput("input/day02.txt")
	if err != nil {
		fmt.Println(err)
		return;
	}
	loopInput(input)
}
