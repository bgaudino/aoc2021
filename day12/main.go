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
	seen godino.Set[*cave]
}

func (c cave) isSmall() bool {
	for _, ch := range c.name {
		if unicode.IsUpper(ch) {
			return false
		}
	}
	return true
}

func main() {
	caves := map[string]*cave{}
	file, _ := os.Open("../data/day12.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "-")
		a, b := s[0], s[1]
		c1 := caves[a]
		c2 := caves[b]
		if c1 == nil {
			c1 = &cave{a, []*cave{}}
		}
		if c2 == nil {
			c2 = &cave{b, []*cave{}}
		}
		c1.neighbors = append(c1.neighbors, c2)
		c2.neighbors = append(c2.neighbors, c1)
		caves[a] = c1
		caves[b] = c2
	}

	start := state{caves["start"], godino.NewSet[*cave]()}
	q := godino.NewDeque[state]()
	q.PushRight(start)
	paths := 0
	for q.Len() > 0 {
		s := q.PopRight()
		if s.cave.name == "end" {
			paths++
			continue
		}
		s.seen.Add(s.cave)
		for _, n := range s.cave.neighbors {
			if !n.isSmall() || !s.seen.Has(n) {
				q.PushRight(state{n, s.seen.Copy()})
			}
		}
	}
	fmt.Println(numPaths(caves))
}

func numPaths(c map[string]*cave) int {
	paths := 0
	start := state{c["start"], godino.NewSet[*cave]()}
	q := godino.NewDeque[state]()
	q.PushRight(start)
	for q.Len() > 0 {
		s := q.PopRight()
		if s.cave.name == "end" {
			paths++
			continue
		}
		s.seen.Add(s.cave)
		for _, n := range s.cave.neighbors {
			if !n.isSmall() || !s.seen.Has(n) {
				q.PushRight(state{n, s.seen.Copy()})
			}
		}
	}
	return paths
}
