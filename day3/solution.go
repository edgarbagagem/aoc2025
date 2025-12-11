package day3

import (
	"edgar/aoc2025/utils"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Solution(part int) error {
	switch part {
	case 1:
		return partOne()
	case 2:
		return partTwo()
	default:
		return fmt.Errorf("invalid part %d for day 3", part)
	}
}

func partOne() error {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}
	inputPath := os.Args[1]

	data, err := utils.InputStrings(inputPath)
	if err != nil {
		log.Fatalf("error parsing data from input: %s", err)
	}

	sum := 0
	for _, line := range data {
		best := 0
		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				val := int(line[i]-'0')*10 + int(line[j]-'0')
				if val > best {
					best = val
				}
			}
		}
		sum += best
	}
	fmt.Println(sum)

	return nil
}

func partTwo() error {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}
	inputPath := os.Args[1]

	data, err := utils.InputStrings(inputPath)
	if err != nil {
		log.Fatalf("error parsing data from input: %s", err)
	}

	sum := int64(0)
	for _, bank := range data {
		k := 12
		drop := len(bank) - k
		stack := make([]byte, 0, len(bank))

		for i := 0; i < len(bank); i++ {
			battery := bank[i]
			for len(stack) > 0 && drop > 0 && stack[len(stack)-1] < battery {
				stack = stack[:len(stack)-1]
				drop--
			}
			stack = append(stack, battery)
		}

		stack = stack[:k]

		joltage, err := strconv.ParseInt(string(stack), 10, 64)
		if err != nil {
			panic("error converting joltage to integer")
		}

		sum += joltage

	}
	fmt.Println(sum)
	return nil
}
