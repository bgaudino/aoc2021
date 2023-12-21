package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/bgaudino/godino"
)

type cave struct {
	name      string
	neighbors []*cave
}

type state struct {
	cave *cave
	seen godino.Counter[*cave]
}

func (c cave) isSmall() bool {
	for _, ch := range c.name {
		if unicode.IsUpper(ch) {
			return false
		}
	}
	return true
}

func copyCounter(m godino.Counter[*cave]) godino.Counter[*cave] {
	elements := []*cave{}
	for _, e := range m.Elements() {
		for i := 0; i < e.Count; i++ {
			elements = append(elements, e.Element)
		}
	}
	return godino.NewCounter(elements)
}

func main() {
	caves := map[string]*cave{}
	file, _ := os.Open("../data/day12.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "-")
		a, b := s[0], s[1]
		c1, ok1 := caves[a]
		c2, ok2 := caves[b]
		if !ok1 {
			c1 = &cave{a, []*cave{}}
		}
		if !ok2 {
			c2 = &cave{b, []*cave{}}
		}
		c1.neighbors = append(c1.neighbors, c2)
		c2.neighbors = append(c2.neighbors, c1)
		caves[a] = c1
		caves[b] = c2
	}

	fmt.Printf("Part 1: %v\n", numPaths(caves, false))
	fmt.Printf("Part 2: %v\n", numPaths(caves, true))
}

func numPaths(c map[string]*cave, part2 bool) int {
	paths := 0
	start := state{c["start"], godino.NewCounter([]*cave{})}
	q := godino.NewDeque[state]()
	q.PushRight(start)
	for q.Len() > 0 {
		s := q.PopRight()
		if s.cave.isSmall() {
			s.seen.Add(s.cave)
		}
		if s.cave.name == "end" {
			paths++
			continue
		}
		for _, n := range s.cave.neighbors {
			if n == c["start"] {
				continue
			}
			if n.isSmall() && s.seen.Get(n) >= 1 {
				if part2 {
					if s.seen.MostCommon(1)[0].Count >= 2 {
						continue
					}
				} else {
					continue
				}
			}
			q.PushRight(state{n, copyCounter(s.seen)})
		}
	}
	return paths
}
