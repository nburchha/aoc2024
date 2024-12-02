package main

import (
	"fmt"
	"aoc2024/utils"
)

func removeElement(slice []int, index int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func checkReport(report []int) bool {
	decreasing, increasing := 1, 1
	for i := 1; i < len(report); i++ {
		if abs(report[i-1] - report[i]) > 3 {
			return false
		} else if report[i] < report[i-1] {
			decreasing++
		} else if (report[i] > report[i-1]) {
			increasing++
		}
	}
	if decreasing == len(report) || increasing == len(report) {
		return true
	}
	return false
}

func checkOptions(report []int) bool {
	if checkReport(report) {
		return true
	}
	//this is for part 2
	for index := range report {
		tmp := removeElement(report, index)
		if checkReport(tmp) {
			return true
		}
	}
	return false
}

func processInput(input [][]int) {
	count := 0
	for _, report := range input {
		if checkOptions(report) {
			count++
		}
	}
	fmt.Println("Safe reports:", count)
}

func main() {
	input, err := utils.ParseInput("input/day02.txt")
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}
	processInput(input)
}
