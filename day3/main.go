package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"time"
)

func main() {
	file, err := os.ReadFile("/home/prathith/Documents/advent_of_code_2025/inputs/day03.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s", err)
	}

	s := string(file)
	s = s[:len(s)-1]

	lines := strings.Split(s, "\n")
	timeStart := time.Now()
	fmt.Printf("Part 1: %d\n", part1(lines))
	timeEnd := time.Now()
	fmt.Println("First part took ", timeEnd.Sub(timeStart))
	fmt.Printf("Part 2: %d\n", part2())
}

func part1(lines []string) int {
	total := 0

	for _, line := range lines {
		tenths := -1
		units := -1
		for i, r := range line {
			n, _ := strconv.Atoi(string(r))

			if n > tenths && i < len(line) - 1 {
				tenths = n
				units  = -1
				continue
			}

			if n > units {
				units = n
			}
		}

		total += tenths*10 + units
	}
	return total
}

func part2() int {
	return 0
}
