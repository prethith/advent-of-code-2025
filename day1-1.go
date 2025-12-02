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

	pos := 50
	count := 0

	scanner := bufio.NewScanner(file);

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

	fmt.Printf("Number of zeros: %d", count)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err);
	}
}
