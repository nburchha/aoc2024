package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[[2]int]int)
var maxDepth = 75

func recIter(val, depth int) int {
	if depth >= maxDepth { return 1 }
	if val == 0 { return recIter(1, depth+1) }
	str := strconv.Itoa(val)
	if len(str) > 1 && len(str) % 2 == 0 {
		str1 := str[:len(str)/2]
		str2 := str[len(str)/2:]
		num1, _ := strconv.Atoi(str1)
		num2, _ := strconv.Atoi(str2)
		if _, err := cache[[2]int{num1, depth+1}]; !err {
			cache[[2]int{num1, depth+1}] = recIter(num1, depth+1)
		}
		if _, err := cache[[2]int{num2, depth+1}]; !err {
			cache[[2]int{num2, depth+1}] = recIter(num2, depth+1)
		}
		return cache[[2]int{num1, depth+1}] + cache[[2]int{num2, depth+1}]
	} else {
		return recIter(val * 2024, depth+1)
	}
}

func main() {
	input, _ := os.ReadFile("input/day11.txt")
	words := strings.Fields(string(input))

	var nums []int
	for _, word := range words {
		num, _ := strconv.Atoi(word)
		nums = append(nums, int(num))
	}

	totalCount := 0

	for _, startNum := range nums {
		totalCount += recIter(int(startNum), 0)
	}

	fmt.Println("Total vals after",maxDepth,"blinks:",totalCount)
}