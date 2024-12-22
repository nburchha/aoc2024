package main
import (
	"fmt"
	"strconv"
	"os"
	"strings"
)

const (
	FILE = "input/day22.txt"
	// FILE = "input/testinput22"
)

func evolveSecretNumber(secretNumber int) int {
	const modulus = 16777216

	// Step 1: Multiply by 64, mix, and prune
	secretNumber ^= secretNumber * 64
	secretNumber %= modulus

	// Step 2: Divide by 32 (integer division), mix, and prune
	secretNumber ^= secretNumber / 32
	secretNumber %= modulus

	// Step 3: Multiply by 2048, mix, and prune
	secretNumber ^= secretNumber * 2048
	secretNumber %= modulus

	return secretNumber
}

func main() {
	var secretNums []int
	input, _ := os.ReadFile(FILE)
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		secretNums = append(secretNums, num)
	}

	for i := 0; i < 2000; i++ {
		for j, num := range secretNums {
			secretNums[j] = evolveSecretNumber(num)
		}
	}
	sum := 0
	for _, num := range secretNums {
		sum += num
	}
	fmt.Println(sum)
}