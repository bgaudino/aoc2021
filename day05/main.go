package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bgaudino/godino"
)

type point struct {
	x int
	y int
}

func (p1 point) line(p2 point, diagonals bool) []point {
	points := []point{}
	xStep, yStep := 1, 1
	if p1.x > p2.x {
		xStep = -1
	}
	if p1.y > p2.y {
		yStep = -1
	}
	if p1.x == p2.x {
		for y := p1.y; y != p2.y+yStep; y += yStep {
			points = append(points, point{p1.x, y})
		}
	} else if p1.y == p2.y {
		for x := p1.x; x != p2.x+xStep; x += xStep {
			points = append(points, point{x, p1.y})
		}
	} else if diagonals {
		p3 := point{p1.x, p1.y}
		for p3 != p2 {
			points = append(points, p3)
			p3.x += xStep
			p3.y += yStep
		}
		points = append(points, p2)
	}
	return points
}

func main() {
	part1 := godino.NewCounter([]point{})
	part2 := godino.NewCounter([]point{})

	file, _ := os.Open("../data/day05.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " -> ")
		points := [2]point{}
		for i, s := range split {
			nums := strings.Split(s, ",")
			x, _ := strconv.Atoi(nums[0])
			y, _ := strconv.Atoi(nums[1])
			points[i] = point{x, y}
		}
		part1.Update(points[0].line(points[1], false))
		part2.Update(points[0].line(points[1], true))
	}
	for i, part := range []godino.Counter[point]{part1, part2} {
		count := 0
		for _, point := range part.MostCommon(-1) {
			if point.Count >= 2 {
				count++
			} else {
				fmt.Printf("Part %v: %v\n", i+1, count)
				break
			}
		}
	}
}
