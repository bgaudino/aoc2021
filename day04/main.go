package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/bgaudino/godino"
)

type square struct {
	x int
	y int
}

type bingoCard struct {
	numbers map[int]square
	matches godino.Set[square]
}

func (b bingoCard) mark(n int) {
	s, ok := b.numbers[n]
	if ok {
		b.matches.Add(s)
	}
}

func (b *bingoCard) isWinner(ways [][]square) bool {
	for _, way := range ways {
		if godino.Every(way, func(s square) bool { return b.matches.Has(s) }) {
			return true
		}
	}
	return false
}

func (b bingoCard) score(n int) (s int) {
	for no, sq := range b.numbers {
		if !b.matches.Has(sq) {
			s += no
		}
	}
	return s * n
}

func main() {
	file, _ := os.Open("../data/day04.txt")
	scanner := bufio.NewScanner(file)

	i := 0
	numbers := []int{}
	card := [][]int{}
	cards := []bingoCard{}
	re := regexp.MustCompile(`[\s,]+`)
	for scanner.Scan() {
		row := []int{}
		text := scanner.Text()
		if text != "" {
			for _, s := range re.Split(text, -1) {
				if s == "" {
					continue
				}
				n, _ := strconv.Atoi(s)
				row = append(row, n)
			}
			if i == 0 {
				numbers = row
			} else {
				card = append(card, row)
			}
		}
		if len(card) == 5 {
			b := bingoCard{make(map[int]square), godino.NewSet[square]()}
			for y, r := range card {
				for x, n := range r {
					b.numbers[n] = square{x, y}
				}
			}
			cards = append(cards, b)
			card = [][]int{}
		}
		i++
	}

	waysToWin := [][]square{}
	for y := 0; y < 5; y++ {
		row := []square{}
		for x := 0; x < 5; x++ {
			row = append(row, square{x, y})
		}
		waysToWin = append(waysToWin, row)
	}
	for x := 0; x < 5; x++ {
		row := []square{}
		for y := 0; y < 5; y++ {
			row = append(row, square{x, y})
		}
		waysToWin = append(waysToWin, row)
	}

	scores := []int{}
	winners := godino.NewSet[int]()
	for _, n := range numbers {
		for j, c := range cards {
			if winners.Has(j) {
				continue
			}
			c.mark(n)
			if c.isWinner(waysToWin) {
				scores = append(scores, c.score(n))
				winners.Add(j)
			}
		}
	}

	fmt.Printf("Part 1: %v\n", scores[0])
	fmt.Printf("Part 2: %v\n", scores[len(scores)-1])
}
