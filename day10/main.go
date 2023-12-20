package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/bgaudino/godino"
)

func main() {
	closingCharacters := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	openingCharacters := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	syntaxErrorPoints := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	completionPoints := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	file, _ := os.Open("../data/day10.txt")
	scanner := bufio.NewScanner(file)
	totalSyntaxErrorScore := 0
	completionScores := []int{}
	for scanner.Scan() {
		stack := godino.NewDeque[rune]()
		isValid := true
		text := scanner.Text()
		for _, c := range text {
			opener, isClosing := closingCharacters[c]
			if isClosing {
				if stack.PeekRight() == opener {
					stack.PopRight()
				} else {
					isValid = false
					totalSyntaxErrorScore += syntaxErrorPoints[c]
					break
				}
			} else {
				stack.PushRight(c)
			}
		}
		if isValid {
			score := 0
			for stack.Len() > 0 {
				score *= 5
				score += completionPoints[openingCharacters[stack.PopRight()]]
			}
			completionScores = append(completionScores, score)
		}
	}
	fmt.Printf("Part 1: %v\n", totalSyntaxErrorScore)
	sort.Ints(completionScores)
	part2 := completionScores[len(completionScores)/2]
	fmt.Printf("Part 2: %v\n", part2)
}
