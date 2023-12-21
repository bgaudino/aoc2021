package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/bgaudino/godino"
)

type coord struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("../data/day11.txt")
	scanner := bufio.NewScanner(file)
	octopuses := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, c := range scanner.Text() {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		octopuses = append(octopuses, row)
	}
	flashes := 0
	numOctopuses := len(octopuses[0]) * len(octopuses)
	i := 1
	for {
		newFlashes := step(octopuses)
		if i < 100 {
			flashes += newFlashes
		}
		if newFlashes == numOctopuses {
			break
		}
		i++
	}
	fmt.Printf("Part 1: %v\n", flashes)
	fmt.Printf("Part 2: %v\n", i)
}

func step(octopuses [][]int) int {
	flashed := godino.NewSet[coord]()
	for y, row := range octopuses {
		for x := range row {
			octopuses[y][x]++
		}
	}
	for {
		didFlash := false
		for y, row := range octopuses {
			for x, o := range row {
				c := coord{x, y}
				if o > 9 && !flashed.Has(c) {
					flash(octopuses, c)
					didFlash = true
					flashed.Add(c)
				}
			}
		}
		if !didFlash {
			break
		}
	}
	for _, c := range flashed.Members() {
		octopuses[c.y][c.x] = 0
	}
	return len(flashed)
}

func flash(octopuses [][]int, o coord) {
	h, w := len(octopuses), len(octopuses[0])
	neighbors := []coord{
		{o.x - 1, o.y - 1}, {o.x, o.y - 1}, {o.x + 1, o.y - 1},
		{o.x - 1, o.y}, {o.x + 1, o.y},
		{o.x - 1, o.y + 1}, {o.x, o.y + 1}, {o.x + 1, o.y + 1},
	}
	for _, n := range neighbors {
		if n.x >= 0 && n.x < w && n.y >= 0 && n.y < h {
			octopuses[n.y][n.x]++
		}
	}
}
