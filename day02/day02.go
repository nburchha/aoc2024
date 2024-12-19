package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInput(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var result [][]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		intLine := make([]int, len(line))
		for i, str := range line {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", str)
				return nil, err
			}
			intLine[i] = num
		}
		result = append(result, intLine)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

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
	input, err := parseInput("input/day02.txt")
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}
	processInput(input)
}
