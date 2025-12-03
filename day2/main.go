package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	data, err := os.ReadFile("/home/prathith/Documents/advent_of_code_2025/inputs/day02.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s", err)
	}
	input := string(data)
	input = strings.TrimSpace(input)
	ranges := strings.Split(input, ",")

	sum := 0

	for _, elem := range ranges {
		elem = strings.TrimSpace(elem)
		index := strings.Index(elem, "-")
		start, err := strconv.Atoi(elem[0:index])
		if err != nil {
			log.Fatalln(err)
		}

		end, err := strconv.Atoi(elem[(index+1):])

		if err != nil {
			log.Fatalln(err)
		}

		for i := start; i <= end; i++ {
			num_str := strconv.Itoa(i)
			if len(num_str) % 2 == 0 {
				sub1 := num_str[0:(len(num_str)/2)]
				sub2 := num_str[(len(num_str)/2):]
				if sub1 == sub2 {
					sum += i
				}
			}
		}
	} 
	return sum
}

func part2() int {
	data, err := os.ReadFile("/home/prathith/Documents/advent_of_code_2025/inputs/day02.txt")
	if err != nil {
		log.Fatalf("Failed to open file %s", err)
	}
	input := string(data)
	input = strings.TrimSpace(input)
	ranges := strings.Split(input, ",")

	sum := 0

	for _, elem := range ranges {
		elem = strings.TrimSpace(elem)
		index := strings.Index(elem, "-")
		start, err := strconv.Atoi(elem[0:index])
		if err != nil {
			log.Fatalln(err)
		}

		end, err := strconv.Atoi(elem[(index+1):])

		if err != nil {
			log.Fatalln(err)
		}

		for i := start; i <= end; i++ {
			num_str := strconv.Itoa(i)
			if repeatedSubstring(num_str) == true {
				sum += i
			}

		} 
	}

	return sum
}

func repeatedSubstring(s string) bool {
	doubled := s + s
	idx := strings.Index(doubled[1:], s)

	if (idx == -1) {
		return false
	}

	return idx+1 != len(s)
}
