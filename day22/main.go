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

func calcBananas(sequence []int, prices [][]int) int {
	var bananas = 0
	for _, price := range prices {
		for index := 0; index < len(price)-4; index++ {
			flag := true
			for i := 0; i < 4; i++ {
				if sequence[i] != price[index+i+1] - price[index+i] {
					flag = false
					break
				}
			}
			if flag {
				bananas += price[index+4]
				break
			}
		}
	}
	return bananas
}

func main() {
	input, _ := os.ReadFile(FILE)
	lines := strings.Split(string(input), "\n")
	var secretNums []int
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		secretNums = append(secretNums, num)
	}
	var prices [][]int
	for j := range secretNums {
		prices = append(prices, []int{})
		for i := 0; i < 2000; i++ {
			prices[j] = append(prices[j], (secretNums[j]%10))
			secretNums[j] = evolveSecretNumber(secretNums[j])
		}
	}
	sum := 0
	for _, num := range secretNums {
		sum += num
	}
	fmt.Println(sum)

	// Part 2
	var bestPrice = 0
	//try out each possible sequence
	for i := -9; i < 10; i++ {
		for j := -9; j < 10; j++ {
			for k := -9; k < 10; k++ {
				for l := -9; l < 10; l++ {
					var sequence = []int{i, j, k, l}
					tmp := calcBananas(sequence, prices)
					if tmp > bestPrice {
						bestPrice = tmp
					}
				}
			}
		}
	}
	fmt.Println(bestPrice)
}