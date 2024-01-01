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
	isX    bool
	isY    bool
	xx     *pair
	yy     *pair
	parent *pair
}

type regularNumber struct {
	n   *pair
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
	if p.isX {
		p.parent.xx = nil
	} else if p.isY {
		p.parent.yy = nil
	}
}

func (p *pair) head() *pair {
	for p.parent != nil {
		p = p.parent
	}
	return p
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
	for _, num := range h.regularNumbers(false) {
		if found {
			return num.n
		}
		if num.n == p && num.pos == "y" {
			found = true
		}
	}
	return nil
}

func (p *pair) prev() *pair {
	h := p.head()
	var prev *pair
	for _, num := range h.regularNumbers(false) {
		if num.n == p {
			return prev
		}
		prev = num.n
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
			if s.n != nil {
				split(s)
			} else {
				break
			}
		}
	}
}

func (p *pair) regularNumbers(print bool) []regularNumber {
	nums := []regularNumber{}
	if p.xx == nil {
		nums = append(nums, regularNumber{p, "x"})
	} else {
		nums = append(nums, p.xx.regularNumbers(false)...)
	}
	if p.yy == nil {
		nums = append(nums, regularNumber{p, "y"})
	} else {
		nums = append(nums, p.yy.regularNumbers(false)...)
	}
	if print {
		for _, num := range nums {
			if num.pos == "x" {
				fmt.Print(num.n.x)
			} else if num.pos == "y" {
				fmt.Print(num.n.y)
			}
		}
		fmt.Println()
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
	p := pair{int(x), int(y), false, false, nil, nil, nil}
	xx, xnum := s[0].(elementSlice)
	yy, ynum := s[1].(elementSlice)
	if xnum {
		xxx := getPair(xx)
		xxx.parent = &p
		xxx.isX = true
		p.xx = xxx
	}
	if ynum {
		yyy := getPair(yy)
		yyy.parent = &p
		yyy.isY = true
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
	for _, num := range n.regularNumbers(false) {
		var v int
		if num.pos == "x" {
			v = num.n.x
		} else if num.pos == "y" {
			v = num.n.y
		}
		if v >= 10 {
			return num
		}
	}
	return regularNumber{nil, ""}
}

func split(n regularNumber) {
	if n.pos == "x" {
		v := n.n.x
		x := v / 2
		y := v - x
		n.n.x = 0
		n.n.xx = &pair{x, y, true, false, nil, nil, n.n}
	} else if n.pos == "y" {
		v := n.n.y
		x := v / 2
		y := v - x
		n.n.y = 0
		n.n.yy = &pair{x, y, false, true, nil, nil, n.n}
	}
}

func add(n1 *pair, n2 *pair) *pair {
	n1.isX = true
	n2.isY = true
	n3 := &pair{0, 0, false, false, n1, n2, nil}
	n1.parent = n3
	n2.parent = n3
	return n3
}
