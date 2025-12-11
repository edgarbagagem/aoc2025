package day3

import (
	"edgar/aoc2025/utils"
	"fmt"
	"log"
	"os"
)

func Solution(part int) error {
	switch part {
	case 1:
		return partOne()
	default:
		return fmt.Errorf("invalid part %d for day 2", part)
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
	return nil
}
