package day4

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
	case 2:
		return partTwo()
	default:
		return fmt.Errorf("invalid part %d for day 4", part)
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

	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	sum := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] != '@' {
				continue
			}

			adjacent := CountAdjacentPaperRolls(dirs, i, j, data)

			if adjacent < 4 {
				sum++
			}
		}
	}

	fmt.Println(sum)

	return nil
}

func partTwo() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: go run main.go <input_file>")
	}
	inputPath := os.Args[1]

	data, err := utils.InputStrings(inputPath)
	if err != nil {
		return fmt.Errorf("error parsing data from input: %s", err)
	}

	removedCount := 0
	for {
		toRemove := make([][2]int, 0)

		for i := 0; i < len(data); i++ {
			for j := 0; j < len(data[i]); j++ {
				if data[i][j] != '@' {
					continue
				}

				adjacent := CountAdjacentPaperRolls(dirs, i, j, data)

				if adjacent < 4 {
					toRemove = append(toRemove, [2]int{i, j})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, pos := range toRemove {
			i, j := pos[0], pos[1]
			row := []rune(data[i])
			row[j] = '.'
			data[i] = string(row)
		}

		removedCount += len(toRemove)
	}

	fmt.Println(removedCount)
	return nil
}

func CountAdjacentPaperRolls(dirs [][2]int, i int, j int, data []string) int {
	adjacent := 0
	for _, d := range dirs {
		ni := i + d[0]
		nj := j + d[1]

		if ni < 0 || ni >= len(data) || nj < 0 || nj >= len(data[0]) {
			continue
		}

		ch := data[ni][nj]
		if ch == '@' {
			adjacent++
		}
	}
	return adjacent
}

var dirs = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}
