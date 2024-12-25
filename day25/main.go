package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	INPUT = "input/day25.txt"
	// INPUT = "input/testinput25"
)

func parseKeyOrLock(lines []string) ([5]int, bool) {
	var res [5]int
	isKey := true
	if lines[0] == "#####" {
		isKey = false
	}
	for _, line := range lines {
		for col, char := range line {
			if char == '#' {
				res[col]++
			}
		}
	}
	for index := range res {
		res[index] -= 1
	}
	fmt.Println(res)
	return res, isKey
}

func parseInput(lines []string) ([][5]int, [][5]int) {
	keys := [][5]int{}
	locks := [][5]int{}
	
	for i:=0;i<len(lines);i++ {
		if lines[i] != "" {
			tmp, isKey := parseKeyOrLock(lines[i:i+7])
			if isKey {
				keys = append(keys, tmp)
			} else {
				locks = append(locks, tmp)
			}
			i+=7
		}
	}
	return keys, locks
} 

func match(key, lock [5]int) bool {
	for index := range key {
		if key[index] + lock[index] > 5 {
			return false
		}
	}
	return true
}

func main() {
	file, _ := os.ReadFile(INPUT)
	lines := strings.Split(string(file), "\n")
	keys, locks := parseInput(lines)

	var count int
	for _, lock := range locks {
		for _, key := range keys {
			if match(key, lock) {
				count++
			}
		}
	}
	fmt.Println(count)
}