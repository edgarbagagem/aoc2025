package day5

import (
	"edgar/aoc2025/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Solution(part int) error {
	switch part {
	case 1:
		return partOne()
	case 2:
		return partTwo()
	default:
		return fmt.Errorf("invalid part %d for day 5", part)
	}
}

type rng struct {
	lo int
	hi int
}

func partOne() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: go run main.go <input_file>")
	}
	inputPath := os.Args[1]

	data, err := utils.InputStrings(inputPath)
	if err != nil {
		return fmt.Errorf("error parsing data from input: %s", err)
	}

	parts, err := breakIntoParts(data)
	if err != nil {
		return fmt.Errorf("error separrating input into ranges and values")
	}

	rangeLines := parts[0]
	valueLines := parts[1]

	ranges := make([]rng, 0, len(rangeLines))

	for _, line := range rangeLines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return fmt.Errorf("invalid range: %s", line)
		}

		var r rng
		fmt.Sscanf(parts[0], "%d", &r.lo)
		fmt.Sscanf(parts[1], "%d", &r.hi)

		ranges = append(ranges, r)
	}

	fresh := 0
	for _, line := range valueLines {
		var id int
		fmt.Sscanf(line, "%d", &id)

		for _, r := range ranges {
			if id >= r.lo && id <= r.hi {
				fresh++
				break
			}
		}
	}

	fmt.Println(fresh)

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

	parts, err := breakIntoParts(data)
	if err != nil {
		return fmt.Errorf("error separating input into ranges")
	}

	rangeLines := parts[0]

	type rng struct {
		lo int
		hi int
	}

	ranges := make([]rng, 0, len(rangeLines))
	for _, line := range rangeLines {
		p := strings.Split(line, "-")
		if len(p) != 2 {
			return fmt.Errorf("invalid range: %s", line)
		}

		var r rng
		fmt.Sscanf(p[0], "%d", &r.lo)
		fmt.Sscanf(p[1], "%d", &r.hi)
		ranges = append(ranges, r)
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].lo < ranges[j].lo
	})

	merged := []rng{ranges[0]}
	for _, r := range ranges[1:] {
		last := &merged[len(merged)-1]

		if r.lo <= last.hi+1 {
			if r.hi > last.hi {
				last.hi = r.hi
			}
		} else {
			merged = append(merged, r)
		}
	}

	total := 0
	for _, r := range merged {
		total += r.hi - r.lo + 1
	}

	fmt.Println(total)
	return nil
}

func breakIntoParts(lines []string) ([][]string, error) {
	var blocks [][]string
	current := []string{}

	for _, line := range lines {
		if line == "" {
			blocks = append(blocks, current)
			current = []string{}
			continue
		}
		current = append(current, line)
	}
	blocks = append(blocks, current)

	return blocks, nil

}
