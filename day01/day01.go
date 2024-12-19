package main

import (
	"bufio"
	"fmt"
	"os"
	"container/heap"
	"strings"
	"strconv"
)

// pair up smallest two, then second smallest two, and so on - add abs(differences) to total

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func parseInput() (*IntHeap, *IntHeap) {
	file, err := os.Open("input/day01-0.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	left := &IntHeap{}
	heap.Init(left)
	right := &IntHeap{}
	heap.Init(right)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := strings.Split(scanner.Text(), "   ")
		num1, err1 := strconv.Atoi(line[0])
		num2, err2 := strconv.Atoi(line[1])
		if err1 != nil || err2 != nil {
			fmt.Println(line, err1, err2)
			fmt.Println("Error parsing input")
			return nil, nil
		}
		heap.Push(left, num1)
		heap.Push(right, num2)
	}
	return left, right
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateDiff(left *IntHeap, right *IntHeap) int {
	total := 0
	for left.Len() > 0 {
		total += abs(heap.Pop(left).(int) - heap.Pop(right).(int))
	}
	return total
}

func calculateSimilarity(left *IntHeap, right *IntHeap) int {
	frequencyMap := make(map[int]int)
	for right.Len() > 0 {
		num := heap.Pop(right).(int)
		frequencyMap[num]++
	}
	total := 0
	for left.Len() > 0 {
		num := heap.Pop(left).(int)
		total = total + frequencyMap[num] * num
	}
	return total
}

func main() {
	left, right := parseInput()
	if left == nil || right == nil {
		fmt.Println("Error parsing input")
		return
	}
	// fmt.Println("first exercise (absolute difference):", calculateDiff(left, right))
	fmt.Println("second exercise (total similarity score):", calculateSimilarity(left, right))
}