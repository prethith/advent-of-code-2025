package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("/home/prathith/Documents/advent_of_code_2025/input-day-1.txt");
	if err != nil {
		log.Fatalf("Failed to open file %s", err);
	}

	defer file.Close()

	dial := 50
	count := 0

	scanner := bufio.NewScanner(file);

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
					count += 1 + (num-dial) / 100
				}

				dial = ((dial - num) % 100 + 100) % 100
			case "R":
				count += (dial + num) / 100
				dial = (dial + num) % 100
		}

		
	}

	fmt.Printf("Number of zeros: %d", count)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err);
	}
}
