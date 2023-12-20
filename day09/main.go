package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/bgaudino/godino"
)

type coord struct {
	x int
	y int
}

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
	basins := []godino.Set[coord]{}
	seen := godino.NewSet[coord]()
	h, w := len(heightMap), len(heightMap[0])
	for y, row := range heightMap {
		for x, height := range row {
			isLow := true
			c := coord{x, y}
			for _, n := range getNeighbors(c, h, w) {
				if height >= heightMap[n.y][n.x] {
					isLow = false
					break
				}
			}
			if isLow {
				riskLevels += height + 1
			}
			if height != 9 && !seen.Has(c) {
				basin := findBasin(c, heightMap)
				basins = append(basins, basin)
				seen.Update(basin)
			}
		}
	}
	fmt.Printf("Part 1: %v\n", riskLevels)

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})
	part2 := godino.Prod(godino.Map(basins[:3], func(b godino.Set[coord]) int { return len(b) })...)
	fmt.Printf("Part 2: %v\n", part2)
}

func getNeighbors(c coord, h int, w int) []coord {
	return godino.Filter(
		[]coord{
			{c.x, c.y - 1},
			{c.x - 1, c.y}, {c.x + 1, c.y},
			{c.x, c.y + 1},
		},
		func(c coord) bool {
			return c.x >= 0 && c.x < w && c.y >= 0 && c.y < h
		},
	)
}

func findBasin(c coord, m [][]int) godino.Set[coord] {
	b := godino.NewSet[coord]()
	h, w := len(m), len(m[0])
	q := godino.NewDeque[coord]()
	q.PushRight(c)
	for q.Len() > 0 {
		location := q.PopLeft()
		b.Add(location)
		for _, n := range getNeighbors(location, h, w) {
			if m[n.y][n.x] != 9 && !b.Has(n) {
				q.PushRight(n)
			}
		}
	}
	return b
}
