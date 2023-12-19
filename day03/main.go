package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/bgaudino/godino"
)

func mostAndLeastCommon(report []string, position int) (byte, byte) {
	c := godino.NewCounter([]byte{})
	for _, number := range report {
		c.Add(number[position])
	}
	sorted := c.MostCommon(-1)
	most, least := sorted[0], sorted[len(sorted)-1]
	if most.Count == least.Count {
		return '1', '0'
	}
	return most.Element, least.Element
}

func main() {
	file, _ := os.Open("../data/day03.txt")
	scanner := bufio.NewScanner(file)
	report := []string{}
	for scanner.Scan() {
		report = append(report, scanner.Text())
	}

	gammaRate := []byte{}
	epsilonRate := []byte{}
	l := len(report[0])
	for i := 0; i < l; i++ {
		most, least := mostAndLeastCommon(report, i)
		gammaRate = append(gammaRate, most)
		epsilonRate = append(epsilonRate, least)
	}

	i := 0
	oxygenReport := report
	for len(oxygenReport) > 1 {
		j := i % l
		most, _ := mostAndLeastCommon(oxygenReport, j)
		oxygenReport = godino.Filter(oxygenReport, func(n string) bool { return n[j] == most })
		i++
	}

	i = 0
	co2report := report
	for len(co2report) > 1 {
		j := i % l
		_, least := mostAndLeastCommon(co2report, j)
		co2report = godino.Filter(co2report, func(n string) bool { return n[j] == least })
		i++
	}

	gammaRateDecimal, _ := strconv.ParseInt(string(gammaRate), 2, 64)
	epsilonRateDecimal, _ := strconv.ParseInt(string(epsilonRate), 2, 64)
	oxygenGeneratorDecimal, _ := strconv.ParseInt(string(oxygenReport[0]), 2, 64)
	co2scrubberDecimal, _ := strconv.ParseInt(string(co2report[0]), 2, 64)

	fmt.Printf("Part 1: %v\n", gammaRateDecimal*epsilonRateDecimal)
	fmt.Printf("Part 2: %v\n", oxygenGeneratorDecimal*co2scrubberDecimal)
}
