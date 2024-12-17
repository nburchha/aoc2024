package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func resolveComboOperand(op int, A, B, C int) int {
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	default:
		panic("Invalid 7")
	}
}

func main() {
	file, err := os.Open("input/day17.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var A, B, C int
	var program []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Register A:") {
			parts := strings.Split(line, ":")
			A, _ = strconv.Atoi(strings.TrimSpace(parts[1]))
		} else if strings.HasPrefix(line, "Register B:") {
			parts := strings.Split(line, ":")
			B, _ = strconv.Atoi(strings.TrimSpace(parts[1]))
		} else if strings.HasPrefix(line, "Register C:") {
			parts := strings.Split(line, ":")
			C, _ = strconv.Atoi(strings.TrimSpace(parts[1]))
		} else if strings.HasPrefix(line, "Program:") {
			parts := strings.Split(line, ":")
			progStr := strings.TrimSpace(parts[1])
			progParts := strings.Split(progStr, ",")
			for _, p := range progParts {
				val, _ := strconv.Atoi(strings.TrimSpace(p))
				program = append(program, val)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ip := 0
	var outputVals []int

	for {
		if ip >= len(program) { break }
		opcode := program[ip]
		if ip+1 >= len(program) { break }
		operand := program[ip+1]

		switch opcode {
		case 0: // adv: A = floor(A / (2^(combo_value)))
			val := resolveComboOperand(operand, A, B, C)
			denominator := int(math.Pow(2, float64(val)))
			A = A / denominator
		case 1: // bxl: B = B XOR literal_operand
			B = B ^ operand
		case 2: // bst: B = (combo_operand_value % 8)
			val := resolveComboOperand(operand, A, B, C)
			B = val % 8
		case 3: // jnz: if A != 0, ip = literal_operand
			if A != 0 {
				ip = operand
				continue
			}
		case 4: // bxc: B = B XOR C (operand ignored)
			B = B ^ C
		case 5: // out: output combo_operand_value % 8
			val := resolveComboOperand(operand, A, B, C)
			outputVals = append(outputVals, val%8)
		case 6: // bdv: B = floor(A / (2^(combo_value)))
			val := resolveComboOperand(operand, A, B, C)
			denominator := int(math.Pow(2, float64(val)))
			B = A / denominator
		case 7: // cdv: C = floor(A / (2^(combo_value)))
			val := resolveComboOperand(operand, A, B, C)
			denominator := int(math.Pow(2, float64(val)))
			C = A / denominator
		default:
			panic("Invalid opcode encountered.")
		}
		ip += 2
	}
	for i, val := range outputVals {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(val)
	}
	fmt.Println()
}