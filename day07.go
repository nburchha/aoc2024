package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type OperatorFunc func(int, int) int

func add(a, b int) int { return a + b }
// func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }
// func divide(a, b int) int { return a / b }
func concat(a, b int) int {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	concatStr := aStr + bStr
	result, _ := strconv.Atoi(concatStr)
	return result
}

var operators = map[int]OperatorFunc{
    0: add,
    1: multiply,
    2: concat,
    // 3: divide,
}

//number: number number ...
//solution: parts which have to be calculated with each other

func parseInput(path string) (map[int][][]int, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    result := make(map[int][][]int)
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ":")
        if len(parts) != 2 {
            continue
        }

        key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
        if err != nil {
            return nil, err
        }

        valueStrings := strings.Fields(parts[1])
        values := make([]int, len(valueStrings))
        for i, v := range valueStrings {
            values[i], err = strconv.Atoi(v)
            if err != nil {
                return nil, err
            }
        }

        // Append to the slice for this key
        result[key] = append(result[key], values)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return result, nil
}

func checkSolution(key int, values []int, operator []int) bool {
	res := values[0]
	for i:=1; i<len(values); i++ {
		res = operators[operator[i - 1]](res, values[i])
	}
	// fmt.Println("Checking solution:", key, values, operator, res)
	return res == key
}

func searchSolution(key int, values []int, operator []int, depth int) bool {
	if depth == len(operator) {
		if checkSolution(key, values, operator) {
			fmt.Println("Found solution:", key, values, operator)
		}
		return checkSolution(key, values, operator)
	}
	for i := 0; i < 3; i++ {
		newOperator := append([]int(nil), operator...)
		newOperator[depth] = i
		if searchSolution(key, values, newOperator, depth+1) {
			return true
		}
	}
	return false
}

func calcSolution(data map[int][][]int) int {
	var count int
	for key, value := range data {
		for _, v := range value {
			operator := make([]int, len(v) - 1)
			if searchSolution(key, v, operator, 0) {
				count += key
			}
		}
	}
	return count
}

func main() {
	data, err := parseInput("input/day07.txt")
	// data, err := parseInput("input/testinput7")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(data)
	fmt.Println(calcSolution(data))
}