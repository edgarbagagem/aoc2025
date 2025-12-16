package day6

import (
	"edgar/aoc2025/utils"
	"fmt"
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
		return fmt.Errorf("invalid part %d for day 6", part)
	}
}

func partOne() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: go run main.go <input_file>")
	}
	inputPath := os.Args[1]

	data, err := utils.InputGrid(inputPath)
	if err != nil {
		return fmt.Errorf("error parsing data from input: %s", err)
	}

	sum := 0

	for i := 0; i < len(data[0]); i++ {
		res, err := strconv.Atoi(data[0][i])
		if err != nil {
			panic(err)
		}

		operator := data[len(data)-1][i]

		for j := 1; j < len(data)-1; j++ {
			value, err := strconv.Atoi(data[j][i])
			if err != nil {
				panic(err)
			}

			switch operator {
			case "*":
				res *= value
			case "+":
				res += value
			}
		}
		sum += res
	}
	fmt.Println(sum)
	return nil
}

func partTwo() error {
	panic("not implemented")
}
