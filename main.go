package main

import (
	"edgar/aoc2025/day1"
	"edgar/aoc2025/day2"
	"edgar/aoc2025/day3"
	"edgar/aoc2025/day4"
	"edgar/aoc2025/day5"
	"edgar/aoc2025/day6"
	"log"
	"os"
	"strconv"
)

func main() {
	argc := len(os.Args)
	if argc < 4 {
		log.Fatal("Usage: go run main.go <input_file> <day_number> <part_number>")
	}

	dayNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("error: ", err)
		return
	}

	if !(dayNum >= 1 && dayNum <= 25) {
		log.Fatal("Invalid day, expected `1 <= arg <= 25`")
		return
	}
	partNum, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal("error: ", err)
		return
	}
	if partNum < 1 || partNum > 2 {
		log.Fatal("Invalid part, expected 1 or 2")
		return
	}

	switch dayNum {
	case 1:
		err = day1.Solution(partNum)
	case 2:
		err = day2.Solution(partNum)
	case 3:
		err = day3.Solution(partNum)
	case 4:
		err = day4.Solution(partNum)
	case 5:
		err = day5.Solution(partNum)
	case 6:
		err = day6.Solution(partNum)
	}

	if err != nil {
		log.Fatal("error: ", err)
	}
}
