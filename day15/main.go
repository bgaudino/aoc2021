package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/bgaudino/godino"
)

type coord struct {
	x int
	y int
}

type state struct {
	risk int
	coord
}

func main() {
	file, _ := os.Open("../data/day15.txt")
	scanner := bufio.NewScanner(file)
	cavern := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, c := range scanner.Text() {
			risk, _ := strconv.Atoi(string(c))
			row = append(row, risk)
		}
		cavern = append(cavern, row)
	}

	fmt.Printf("Part 1: %v\n", safestPath(cavern))
}

func safestPath(cavern [][]int) int {
	height, width := len(cavern), len(cavern[0])
	q := godino.NewDeque[state]()
	q.PushRight(state{0, coord{0, 0}})
	seen := map[coord]int{}
	end := coord{width - 1, height - 1}
	minRisk := int(math.Inf(0))
	for q.Len() > 0 {
		s := q.PopRight()
		seen[s.coord] = s.risk
		if s.coord == end {
			minRisk, _ = godino.Min(minRisk, s.risk)
			continue
		}
		for _, n := range []coord{{s.x, s.y - 1}, {s.x - 1, s.y}, {s.x + 1, s.y}, {s.x, s.y + 1}} {
			if n.x >= 0 && n.x < width && n.y >= 0 && n.y < height {
				newRisk := s.risk + cavern[n.y][n.x]
				risk, ok := seen[n]
				if (!ok || newRisk < risk) && newRisk < minRisk {
					q.PushRight(state{newRisk, n})
				}
			}
		}
	}
	return minRisk
}
