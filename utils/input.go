package utils

import (
	"bufio"
	"os"
	"strings"
)

func InputStrings(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func InputGrid(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var grid [][]string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		grid = append(grid, strings.Fields(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}
