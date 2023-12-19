package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bgaudino/godino"
)

func count(f godino.Dict[int, int]) int {
	return godino.Sum(f.Values()...)
}

func main() {
	file, _ := os.Open("../data/day06.txt")
	scanner := bufio.NewScanner(file)
	fish := make(godino.Dict[int, int])
	for scanner.Scan() {
		text := scanner.Text()
		for _, s := range strings.Split(text, ",") {
			n, _ := strconv.Atoi(s)
			fish[n]++
		}
	}
	for i := 0; i < 256; i++ {
		breeding := fish[0]
		for j := 0; j < 8; j++ {
			fish[j] = fish[j+1]
		}
		fish[6] += breeding
		fish[8] = breeding
		if i == 79 {
			fmt.Printf("Part 1: %v\n", count(fish))
		}
	}
	fmt.Printf("Part 2: %v\n", count(fish))
}
