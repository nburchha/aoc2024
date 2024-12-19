package main

import (
	"log"
	"slices"
)

type Registers struct {
	A int
	B int
	C int
}

func (r *Registers) getComboOperand(index int) int {
	if index < 4 {
		return index
	}
	if index == 4 {
		return r.A
	}
	if index == 5 {
		return r.B
	}
	if index == 6 {
		return r.C
	}
	return 0
}

func processor(program []int, registers *Registers) []int {
	// t := registers.A
	output := make([]int, 0)
	for pc := 0; pc < len(program); {
		instruction := program[pc]
		operand := program[pc+1]
		combo := registers.getComboOperand(operand)
		pc = pc + 2
		switch instruction {
		case 0: // adv: Divide A by 2^combo
			registers.A >>= combo
		case 1: // bxl: B XOR with operand
			registers.B ^= operand
		case 2: // bst: B becomes combo % 8
			registers.B = combo % 8
		case 3: // jnz: Jump if A != 0
			if registers.A != 0 {
				pc = operand
				continue
			}
		case 4: // bxc: B XOR C
			registers.B ^= registers.C
		case 5: // out: Output combo % 8
			output = append(output, combo%8)
		case 6: // bdv: B becomes A divided by 2^combo
			registers.B = (registers.A >> combo)
		case 7: // cdv: C becomes A divided by 2^combo
			registers.C = (registers.A >> combo)
		}
	}
	return output
}

func evaluate(a, b, c int, program []int) {
	log.Println()
	registers := Registers{
		A: a,
		B: b,
		C: c,
	}
	output := processor(program, &registers)
	log.Println("Output:", output)
}

func part2(b, c int, program []int) {
	a := 0
	for i := len(program) - 1; i >= 0; i-- {
		a <<= 3
		registers := Registers{
			A: a,
			B: b,
			C: c,
		}
		for !slices.Equal(processor(program, &registers), program[i:]) {
			a++
			registers = Registers{
				A: a,
				B: b,
				C: c,
			}
		}
	}
	log.Println("Smallest A:", a)
}

func main() {
	evaluate(62769524, 0, 0, []int{2, 4, 1, 7, 7, 5, 0, 3, 4, 0, 1, 7, 5, 5, 3, 0})
	part2(0, 0, []int{2, 4, 1, 7, 7, 5, 0, 3, 4, 0, 1, 7, 5, 5, 3, 0})
}