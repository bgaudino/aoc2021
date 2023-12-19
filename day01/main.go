package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/bgaudino/godino"
)

func main() {
	file, _ := os.Open("../data/day01.txt")
	scanner := bufio.NewScanner(file)
	var depth int
	i := 0
	increases := 0
	windowA := []int{}
	windowB := []int{}
	windowIncreases := 0
	for scanner.Scan() {
		newDepth, _ := strconv.Atoi(scanner.Text())
		if i > 0 && newDepth > depth {
			increases++
		}
		windowA = append(windowA, newDepth)
		if i > 0 {
			windowB = append(windowB, depth)
		}
		if len(windowA) == 4 {
			windowA = windowA[1:]
		}
		if len(windowB) == 4 {
			windowB = windowB[1:]
		}
		if i > 2 && godino.Sum(windowA...) > godino.Sum(windowB...) {
			windowIncreases++
		}
		depth = newDepth
		i++
	}
	fmt.Printf("Part 1: %v\n", increases)
	fmt.Printf("Part 2: %v\n", windowIncreases)
}
