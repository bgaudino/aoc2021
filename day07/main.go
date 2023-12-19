package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/bgaudino/godino"
)

func main() {
	file, _ := os.Open("../data/day07.txt")
	scanner := bufio.NewScanner(file)
	crabs := []int{}
	for scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), ",") {
			n, _ := strconv.Atoi(s)
			crabs = append(crabs, n)
		}
	}
	minPos, _ := godino.Min(crabs...)
	maxPos, _ := godino.Max(crabs...)
	costsA := []int{}
	costsB := []int{}
	for i := minPos; i <= maxPos; i++ {
		a, b := 0, 0
		for _, crab := range crabs {
			d := int(math.Abs(float64(crab) - float64(i)))
			a += d
			b += d * (d + 1) / 2 // triangular number
		}
		costsA = append(costsA, a)
		costsB = append(costsB, b)
	}
	part1, _ := godino.Min(costsA...)
	fmt.Printf("Part 1: %v\n", part1)
	part2, _ := godino.Min(costsB...)
	fmt.Printf("Part 2: %v\n", part2)
}
