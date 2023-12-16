package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../data/day02.txt")
	scanner := bufio.NewScanner(file)
	horizontalPosition := 0
	depth := 0
	aim := 0
	depth2 := 0
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		movement := split[0]
		distance, _ := strconv.Atoi(split[1])
		if movement == "forward" {
			horizontalPosition += distance
			depth2 += aim * distance
		} else if movement == "down" {
			depth += distance
			aim += distance
		} else if movement == "up" {
			depth -= distance
			aim -= distance
		}
	}
	part1 := horizontalPosition * depth
	fmt.Printf("Part 1: %v\n", part1)
	part2 := horizontalPosition * depth2
	fmt.Printf("Part 2: %v\n", part2)
}
