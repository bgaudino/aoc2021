package main

import (
	"fmt"

	"github.com/bgaudino/godino"
)

var X int = 34
var XX int = 67
var Y int = -215
var YY int = -186

type coord struct {
	x int
	y int
}

type probe struct {
	position coord
	velocity coord
	visited  godino.Set[coord]
}

func (p *probe) move() {
	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
	p.visited.Add(p.position)
	if p.velocity.x > 0 {
		p.velocity.x--
	} else if p.velocity.x < 0 {
		p.velocity.x++
	}
	p.velocity.y--
}

func (p probe) fire() (bool, int) {
	for {
		p.move()
		if inTargetArea(p.position) {
			my, _ := godino.Max(godino.Map(p.visited.Members(), func(c coord) int { return c.y })...)
			return true, my
		}
		if missed(p) {
			return false, -1
		}
	}
}

func inTargetArea(c coord) bool {
	return c.x >= X && c.x <= XX && c.y >= Y && c.y <= YY
}

func missed(p probe) bool {
	if p.position.y < YY {
		return true
	}
	if p.velocity.x > 0 && p.position.x > XX {
		return true
	}
	if p.velocity.x < 0 && p.position.x < X {
		return true
	}
	return false
}

func main() {
	start := coord{0, 0}
	yPositions := []int{}
	velocities := godino.NewSet[coord]()
	for x := 0; x <= XX; x++ {
		for y := 500; y >= Y; y-- {
			p := probe{start, coord{x, y}, godino.NewSet(start)}
			hit, maxY := p.fire()
			if hit {
				yPositions = append(yPositions, maxY)
				velocities.Add(p.velocity)
			}
		}
	}
	part1, _ := godino.Max(yPositions...)
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", len(velocities))
}
