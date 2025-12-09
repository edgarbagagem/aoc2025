package day1

import (
	"edgar/aoc2025/utils"
	"fmt"
	"log"
	"math"
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
		return fmt.Errorf("invalid part %d for day 1", part)
	}
}

func partOne() error {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}
	inputPath := os.Args[1]

	data, err := utils.InputStrings(inputPath)
	if err != nil {
		return fmt.Errorf("error parsing data from input: %s", err)
	}

	dial := 50
	zeroCount := 0
	for i, ins := range data {
		turn, err := strconv.Atoi(ins[1:])
		if err != nil {
			return fmt.Errorf("error parsing rotation instruction number on line %d: %q, error: %v", i+1, ins, err)
		}

		if ins[0] == 'L' {
			dial -= turn
		}

		if ins[0] == 'R' {
			dial += turn
		}

		if dial%100 == 0 {
			zeroCount++
		}
	}

	fmt.Println(zeroCount)
	return nil
}

func partTwo() error {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}
	inputPath := os.Args[1]

	data, err := utils.InputStrings(inputPath)
	if err != nil {
		return fmt.Errorf("error parsing data from input: %s", err)
	}

	dial := 50
	zeroCount := 0
	for _, ins := range data {
		turn, err := strconv.Atoi(ins[1:])
		if err != nil {
			panic(err)
		}

		if ins[0] == 'L' {
			turn *= -1
		}

		zeroCount += int(math.Floor(math.Abs(float64(turn)) / 100))
		remainder := turn % 100

		if remainder == 0 {
			continue
		}

		if remainder >= 0 {
			if dial+remainder >= 100 {
				zeroCount += 1
			}
			dial = mod(dial+remainder, 100)
			continue
		}

		if dial+remainder <= 0 && dial != 0 {
			zeroCount += 1
		}
		dial = mod(dial+remainder, 100)
	}

	fmt.Println(zeroCount)
	return nil
}

func mod(a, b int) int {
	return (a%b + b) % b
}
