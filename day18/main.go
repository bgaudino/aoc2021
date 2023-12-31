package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bgaudino/godino"
)

type elementSlice = []interface{}

type pair struct {
	x      int
	y      int
	xx     *pair
	yy     *pair
	parent *pair
}

type regularNumber struct {
	p   *pair
	pos string
}

func (p *pair) explode() {
	prev := p.prev()
	if prev != nil {
		if prev.yy == nil {
			prev.y += p.x
		} else if prev.xx == nil {
			prev.x += p.x
		}
	}
	next := p.next()
	if next != nil {
		if next.xx == nil {
			next.x += p.y
		} else if next.yy == nil {
			next.y += p.y
		}
	}
	if p.isX() {
		p.parent.xx = nil
	} else if p.isY() {
		p.parent.yy = nil
	}
}

func (p *pair) head() *pair {
	for p.parent != nil {
		p = p.parent
	}
	return p
}

func (p *pair) isX() bool {
	return p.parent != nil && p.parent.xx == p
}

func (p *pair) isY() bool {
	return p.parent != nil && p.parent.yy == p
}

func (p *pair) magnitude() int {
	var x int
	if p.xx == nil {
		x = p.x
	} else {
		x = p.xx.magnitude()
	}
	var y int
	if p.yy == nil {
		y = p.y
	} else {
		y = p.yy.magnitude()
	}
	return 3*x + 2*y
}

func (p *pair) next() *pair {
	h := p.head()
	found := false
	for _, num := range h.regularNumbers() {
		if found {
			return num.p
		}
		if num.p == p && num.pos == "y" {
			found = true
		}
	}
	return nil
}

func (p *pair) prev() *pair {
	h := p.head()
	var prev *pair
	for _, num := range h.regularNumbers() {
		if num.p == p {
			return prev
		}
		prev = num.p
	}
	return nil
}

func (p *pair) reduce() {
	for {
		e := shouldExplode(p)
		if e != nil {
			e.explode()
		} else {
			s := shouldSplit(p)
			if s.p != nil {
				split(s)
			} else {
				break
			}
		}
	}
}

func (p *pair) regularNumbers() []regularNumber {
	nums := []regularNumber{}
	if p.xx == nil {
		nums = append(nums, regularNumber{p, "x"})
	} else {
		nums = append(nums, p.xx.regularNumbers()...)
	}
	if p.yy == nil {
		nums = append(nums, regularNumber{p, "y"})
	} else {
		nums = append(nums, p.yy.regularNumbers()...)
	}
	return nums
}

func main() {
	file, _ := os.Open("../data/day18.txt")
	scanner := bufio.NewScanner(file)
	var p *pair
	es := []elementSlice{}
	for scanner.Scan() {
		data := []byte(scanner.Text())
		var s elementSlice
		json.Unmarshal(data, &s)
		es = append(es, s)
		if p == nil {
			p = getPair(s)
		} else {
			p = add(p, getPair(s))
			p.reduce()
		}
	}
	fmt.Printf("Part 1: %v\n", p.magnitude())
	maxMagnitude := 0
	for c := range godino.Permutations(es, 2) {
		p := add(getPair(c[0]), getPair(c[1]))
		p.reduce()
		maxMagnitude, _ = godino.Max(maxMagnitude, p.magnitude())
	}
	fmt.Printf("Part 2: %v\n", maxMagnitude)
}

func getPair(s elementSlice) *pair {
	x, _ := s[0].(float64)
	y, _ := s[1].(float64)
	p := pair{int(x), int(y), nil, nil, nil}
	xx, xnum := s[0].(elementSlice)
	yy, ynum := s[1].(elementSlice)
	if xnum {
		xxx := getPair(xx)
		xxx.parent = &p
		p.xx = xxx
	}
	if ynum {
		yyy := getPair(yy)
		yyy.parent = &p
		p.yy = yyy
	}
	return &p
}

type state struct {
	n *pair
	d int
}

func shouldExplode(n *pair) *pair {
	dq := godino.NewDeque[state]()
	dq.PushRight(state{n, 0})
	for dq.Len() > 0 {
		s := dq.PopLeft()
		if s.d == 4 {
			return s.n
		}
		if s.n.xx != nil {
			dq.PushRight(state{s.n.xx, s.d + 1})
		}
		if s.n.yy != nil {
			dq.PushRight(state{s.n.yy, s.d + 1})
		}
	}
	return nil
}

func shouldSplit(n *pair) regularNumber {
	for _, num := range n.regularNumbers() {
		var v int
		if num.pos == "x" {
			v = num.p.x
		} else if num.pos == "y" {
			v = num.p.y
		}
		if v >= 10 {
			return num
		}
	}
	return regularNumber{nil, ""}
}

func split(n regularNumber) {
	if n.pos == "x" {
		v := n.p.x
		x := v / 2
		y := v - x
		n.p.x = 0
		n.p.xx = &pair{x, y, nil, nil, n.p}
	} else if n.pos == "y" {
		v := n.p.y
		x := v / 2
		y := v - x
		n.p.y = 0
		n.p.yy = &pair{x, y, nil, nil, n.p}
	}
}

func add(n1 *pair, n2 *pair) *pair {
	n3 := &pair{0, 0, n1, n2, nil}
	n1.parent = n3
	n2.parent = n3
	return n3
}
