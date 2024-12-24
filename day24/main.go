package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

const (
	FILE = "input/day24.txt"
	// FILE = "input/testinput24"
)

type Gate struct {
	Op      string
	Input1  string
	Input2  string
	Output  string
}

func parseInput(lines []string) (map[string]int, []Gate) {
	wireValues := make(map[string]int)
	var gates []Gate

	// Parse initial wire values
	for _, line := range lines {
		if strings.Contains(line, "->") {
			parts := strings.Split(line, " -> ")
			operands := strings.Fields(parts[0])

			// Determine gate type and input wires
			if len(operands) == 3 {
				gates = append(gates, Gate{
					Op:     operands[1],
					Input1: operands[0],
					Input2: operands[2],
					Output: parts[1],
				})
			} else {
				gates = append(gates, Gate{
					Op:     "ASSIGN",
					Input1: operands[0],
					Output: parts[1],
				})
			}
		} else {
			// Parse initial wire value
			parts := strings.Split(line, ": ")
			value, _ := strconv.Atoi(parts[1])
			wireValues[parts[0]] = value
		}
	}
	return wireValues, gates
}

func simulateCircuit(wireValues map[string]int, gates []Gate) map[string]int {
	// A map to store wire values
	wireMap := make(map[string]int)
	for k, v := range wireValues {
		wireMap[k] = v
	}

	// Evaluate gates
	for len(gates) > 0 {
		var remaining []Gate
		for _, gate := range gates {
			val1, ok1 := wireMap[gate.Input1]
			val2, ok2 := wireMap[gate.Input2]
			switch gate.Op {
			case "AND":
				if ok1 && ok2 {
					wireMap[gate.Output] = val1 & val2
				} else {
					remaining = append(remaining, gate)
				}
			case "OR":
				if ok1 && ok2 {
					wireMap[gate.Output] = val1 | val2
				} else {
					remaining = append(remaining, gate)
				}
			case "XOR":
				if ok1 && ok2 {
					wireMap[gate.Output] = val1 ^ val2
				} else {
					remaining = append(remaining, gate)
				}
			case "ASSIGN":
				if ok1 {
					wireMap[gate.Output] = val1
				} else {
					remaining = append(remaining, gate)
				}
			}
		}
		gates = remaining
	}
	return wireMap
}

func calculateOutput(wireMap map[string]int) int {
	var zWires []string

	// Collect and sort wires starting with 'z' in ascending order
	for name := range wireMap {
		if strings.HasPrefix(name, "z") {
			zWires = append(zWires, name)
		}
	}
	sort.Strings(zWires)

	// Debug: Print sorted z wires and their values
	fmt.Println("zWires and their values:")
	for _, wire := range zWires {
		fmt.Printf("%s: %d\n", wire, wireMap[wire])
	}

	// Build binary output
	var binaryOutput strings.Builder
	for i := len(zWires) - 1; i >= 0; i-- {
		binaryOutput.WriteString(strconv.Itoa(wireMap[zWires[i]]))
	}

	// Convert binary string to decimal
	binaryString := binaryOutput.String()
	fmt.Println("Binary Output:", binaryString)

	result, _ := strconv.ParseInt(binaryString, 2, 64)
	return int(result)
}


func printWireValues(wireMap map[string]int) {
	// Collect and sort wire names
	var wireNames []string
	for name := range wireMap {
		wireNames = append(wireNames, name)
	}
	sort.Strings(wireNames)

	// Print wire values
	for _, name := range wireNames {
		fmt.Printf("%s: %d\n", name, wireMap[name])
	}
}

func main() {
	input, _ := os.ReadFile(FILE)
	lines := strings.Split(string(input), "\n")
	wireValues, gates := parseInput(lines)
	fmt.Println("Wire values:", wireValues)
	fmt.Println("Gates:", gates)
	wireMap := simulateCircuit(wireValues, gates)
	printWireValues(wireMap)
	output := calculateOutput(wireMap)
	fmt.Println("Output:", output)
}