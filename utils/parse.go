package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(path string) ([][]int, error) {
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