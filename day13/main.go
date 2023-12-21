package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bgaudino/godino"
)

type coord struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("../data/day13.txt")
	scanner := bufio.NewScanner(file)
	dotsDone := false
	dots := godino.NewSet[coord]()
	folds := []coord{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			dotsDone = true
		} else if dotsDone {
			fold := strings.Split(text, " ")[2]
			sp := strings.Split(fold, "=")
			axis := sp[0]
			n, _ := strconv.Atoi(sp[1])
			c := coord{0, 0}
			if axis == "x" {
				c.x = n
			} else if axis == "y" {
				c.y = n
			}
			folds = append(folds, c)
		} else {
			digits := strings.Split(text, ",")
			x, _ := strconv.Atoi(digits[0])
			y, _ := strconv.Atoi(digits[1])
			dots.Add(coord{x, y})
		}
	}
	part1 := 0
	for i, f := range folds {
		fold(dots, f)
		if i == 0 {
			part1 = len(folds)
		}
	}
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Println("Part 2:")
	printMap(dots)
}

func getBounds(dots godino.Set[coord]) (coord, coord) {
	minX, _ := godino.Min(godino.Map(dots.Members(), func(c coord) int { return c.x })...)
	maxX, _ := godino.Max(godino.Map(dots.Members(), func(c coord) int { return c.x })...)
	minY, _ := godino.Min(godino.Map(dots.Members(), func(c coord) int { return c.y })...)
	maxY, _ := godino.Max(godino.Map(dots.Members(), func(c coord) int { return c.y })...)
	return coord{minX, minY}, coord{maxX, maxY}
}

func printMap(dots godino.Set[coord]) {
	min, max := getBounds(dots)
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if dots.Has(coord{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(",")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func fold(dots godino.Set[coord], instruction coord) {
	for _, d := range dots.Members() {
		if instruction.y != 0 && d.y > instruction.y {
			dots.Remove(d)
			d.y = instruction.y - (d.y - instruction.y)
			dots.Add(d)
		} else if instruction.x != 0 && d.x > instruction.x {
			dots.Remove(d)
			d.x = instruction.x - (d.x - instruction.x)
			dots.Add(d)
		}
	}
}
