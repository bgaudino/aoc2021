package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/bgaudino/godino"
)

func main() {
	file, _ := os.Open("../data/day09.txt")
	scanner := bufio.NewScanner(file)
	heightMap := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, c := range scanner.Text() {
			h, _ := strconv.Atoi(string(c))
			row = append(row, h)
		}
		heightMap = append(heightMap, row)
	}

	riskLevels := 0

	maxX := len(heightMap[0])
	maxY := len(heightMap)
	neighbors := func(x int, y int) [][]int {
		return godino.Filter(
			[][]int{
				{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
				{x - 1, y}, {x + 1, y},
				{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
			},
			func(coord []int) bool {
				x, y := coord[0], coord[1]
				return x >= 0 && x < maxX && y >= 0 && y < maxY
			},
		)
	}

	for y, row := range heightMap {
		for x, height := range row {
			isLow := true
			for _, n := range neighbors(x, y) {
				if height >= heightMap[n[1]][n[0]] {
					isLow = false
					break
				}
			}
			if isLow {
				riskLevels += height + 1
			}
		}
	}
	fmt.Printf("Part 1: %v\n", riskLevels)
}
