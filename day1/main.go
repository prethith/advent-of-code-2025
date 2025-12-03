package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

// ------------------ PART 1 ------------------

func part1() int {
	file, err := os.Open("/home/prathith/Documents/advent_of_code_2025/inputs/day01.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s", err)
	}
	defer file.Close()

	pos := 50
	count := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		letter := string(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		// logic
		if letter == "L" {
			pos = (pos - num) % 100
		} else if letter == "R" {
			pos = (pos + num) % 100
		}

		if pos == 0 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	return count
}

// ------------------ PART 2 ------------------

func part2() int {
	file, err := os.Open("/home/prathith/Documents/advent_of_code_2025/inputs/day01.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s", err)
	}
	defer file.Close()

	dial := 50
	count := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		letter := string(line[0])
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch letter {
		case "L":
			if dial == 0 {
				count += num / 100
			} else if num >= dial {
				count += 1 + (num-dial)/100
			}

			dial = ((dial - num) % 100 + 100) % 100

		case "R":
			count += (dial + num) / 100
			dial = (dial + num) % 100
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	return count
}
