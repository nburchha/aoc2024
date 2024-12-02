package main

import (
	"fmt"
	"aoc2024/utils"
)

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func safeToRemove(level int, decreasing int, increasing int, priorLevel int) (bool) {
	// if its decreasing but its increasing, or other way around -> true
	if (decreasing > 1 && level > priorLevel) || (increasing > 1 && level < priorLevel) {
		return true
	} else if (level - priorLevel > 3 || priorLevel - level > 3) {
		return true
	} else if (level == priorLevel) {
		return true
	}
	return false
}

//PART 1
// rules:
// - either all decreasing or all increasing
// - any two adjacent levels differ by at least one and at most three -> no bigger steps than 3, 
func calcSafeOrUnsafe(report []int) (bool){
	increasing, decreasing := 1, 1
	safeDistance := true
	removed := false
	if report[0] - report[1] > 3 || report[1] - report[0] > 3 || report[0] == report[1] {
		// fmt.Println("removed first element:", report[0])
		report = removeElement(report, 0)
		removed = true
	}
	for i:=1; i<len(report); i++ {
		priorLevel := report[i-1]
		level := report[i]
		// fmt.Println("i:", i, "priorLevel:", priorLevel, "level:", level)
		// check if we should move on / skip this one
		if (safeToRemove(level, decreasing, increasing, priorLevel) && !removed) {
			removed = true
			report = removeElement(report, i)
			// fmt.Println("removed:", level)
			if i >= len(report) {
				break
			}
			i--
			level = report[i]
		}
		if level < priorLevel {
			decreasing += 1
		} else if level > priorLevel {
			increasing += 1
		}
		if priorLevel - level > 3 || level - priorLevel > 3 || priorLevel == level {
			safeDistance = false
		}
	}
	if (increasing == len(report) || decreasing == len(report)) && safeDistance {
		fmt.Println("Safe")
		return true
	} else {
		fmt.Println("Unsafe")
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
