package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bgaudino/godino"
)

type pair struct {
	e1 byte
	e2 byte
}

func main() {
	file, _ := os.Open("../data/day14.txt")
	scanner := bufio.NewScanner(file)
	var template []byte
	parsedTemplate := false
	insertions := map[pair]byte{}
	for scanner.Scan() {
		b := scanner.Bytes()
		if len(b) == 0 {
			continue
		} else if parsedTemplate {
			insertions[pair{b[0], b[1]}] = b[6]
		} else {
			template = b
			parsedTemplate = true
		}
	}
	for i := 0; i < 10; i++ {
		template = step(template, insertions)
	}
	c := godino.NewCounter(template)
	mc := c.MostCommon(-1)
	part1 := mc[0].Count - mc[len(mc)-1].Count
	fmt.Printf("Part 1: %v\n", part1)
}

func step(template []byte, insertions map[pair]byte) []byte {
	result := []byte{template[0]}
	for i := 1; i < len(template); i++ {
		b, ok := insertions[pair{template[i-1], template[i]}]
		if ok {
			result = append(result, b)
		}
		result = append(result, template[i])
	}
	return result
}
