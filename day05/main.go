package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type Pair struct {
    X int
    Y int
}

func ParseInput(path string) ([]Pair, [][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var pairs []Pair
	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Println("Empty line found")
			break
		}
		line := strings.Split(scanner.Text(), "|")
		if len(line) < 2 {
			continue
		}
		x, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println("Error converting string to int:", line[0])
			return nil, nil, err
		}
		y, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println("Error converting string to int:", line[1])
			return nil, nil, err
		}
		pair := Pair{X: x, Y: y}
		pairs = append(pairs, pair)
	}

	var list [][]int
	i := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		intLine := make([]int, len(line))
		for i, str := range line {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:", str)
				return nil, nil, err
			}
			intLine[i] = num
		}
		list = append(list, intLine)
		i += 1
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return pairs, list, nil
}

// func insertAt(slice []int, index int, values ...int) []int {
//     if index < 0 || index > len(slice) {
//         panic("index out of range")
//     }
//     return append(slice[:index], append(values, slice[index:]...)...)
// }

func swapWithIndex(slice []int, index1, index2 int) {
    if index1 >= 0 && index1 < len(slice) && index2 >= 0 && index2 < len(slice) {
        slice[index1], slice[index2] = slice[index2], slice[index1]
    }
}

func checkNum(pair Pair, pairs []Pair) bool {
	for _, p := range pairs {
		if p.X == pair.Y && p.Y == pair.X {
			return false
		}
	}
	return true
}

func validLine(line []int, pairs []Pair) bool {
	for i, num := range line {
		for j:=0;j<i;j++ {
			if !checkNum(Pair{line[j], num}, pairs) {
				return false
			}
		}
	}
	return true
}

func checkLine(line []int, pairs []Pair) int {
    for {
        if validLine(line, pairs) {
            break
        }
        for i, num := range line {
            valid := true
            for j := 0; j < i; j++ {
                if !checkNum(Pair{line[j], num}, pairs) {
                    valid = false
                    swapWithIndex(line, i, j)
                }
            }
            if valid {
                line = append(line, num)
            }
        }
    }
    // Return the middle number of the line
    middleIndex := len(line) / 2
    return line[middleIndex]
}

func iterateMatrix(matrix [][]int, pairs []Pair) int {
	count := 0
	for _, line := range matrix {
		count = checkLine(line, pairs)
	}
	return count
}

func main() {
    pairs, matrix, err := ParseInput("input/testinput5.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    for _, pair := range pairs {
        fmt.Printf("X: %d, Y: %d\n", pair.X, pair.Y)
    }
	for _, line := range matrix {
		fmt.Println(line)
	}

	//iterate over each list in matrix and check
	fmt.Println(iterateMatrix(matrix, pairs))
}