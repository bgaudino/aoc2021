package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bgaudino/godino"
)

type packet struct {
	version       int64
	typeID        int64
	literalValues []string
	subpackets    []packet
}

func (p packet) LiteralValue() int64 {
	return binaryToDecimal(strings.Join(p.literalValues, ""))
}

func (p packet) VersionSum() int64 {
	s := p.version
	for _, sp := range p.subpackets {
		s += sp.VersionSum()
	}
	return s
}

func (p packet) Value() int64 {
	if p.typeID == 4 {
		return p.LiteralValue()
	}
	subValues := godino.Map(p.subpackets, func(sb packet) int64 { return sb.Value() })
	switch p.typeID {
	case 0:
		return godino.Sum(subValues...)
	case 1:
		return godino.Prod(subValues...)
	case 2:
		m, _ := godino.Min(subValues...)
		return m
	case 3:
		m, _ := godino.Max(subValues...)
		return m
	case 5:
		if subValues[0] > subValues[1] {
			return 1
		} else {
			return 0
		}
	case 6:
		if subValues[0] < subValues[1] {
			return 1
		} else {
			return 0
		}
	case 7:
		if subValues[0] == subValues[1] {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

func binaryToDecimal(binary string) int64 {
	decimal, _ := strconv.ParseInt(binary, 2, 64)
	return decimal
}

var hexToBinary map[rune]string = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func hexadecimalToBinary(hex string) string {
	b := []string{}
	for _, c := range hex {
		b = append(b, hexToBinary[c])
	}
	return strings.Join(b, "")
}

func parsePacket(text string) (packet, string) {
	p := packet{
		version:       binaryToDecimal(text[:3]),
		typeID:        binaryToDecimal(text[3:6]),
		literalValues: []string{},
		subpackets:    []packet{},
	}
	left := ""
	if p.typeID == 4 {
		i := 6
		for {
			p.literalValues = append(p.literalValues, text[i+1:i+5])
			if text[i] == '0' {
				break
			}
			i += 5
		}
		left = text[i+5:]
	} else if text[6] == '0' {
		subpacketLength := binaryToDecimal(text[7 : 7+15])
		left = text[7+15:]
		for subpacketLength > 0 {
			var subpacket packet
			var l string
			subpacket, l = parsePacket(left)
			p.subpackets = append(p.subpackets, subpacket)
			subpacketLength -= int64(len(left) - len(l))
			left = l
		}
	} else {
		numSubpackets := binaryToDecimal(text[7 : 7+11])
		left = text[7+11:]
		for i := 0; i < int(numSubpackets); i++ {
			var subpacket packet
			subpacket, left = parsePacket(left)
			p.subpackets = append(p.subpackets, subpacket)
		}
	}
	return p, left
}

func main() {
	file, _ := os.Open("../data/day16.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("Part 1: %v\n", part1(text))
		fmt.Printf("Part 2: %v\n", part2(text))
	}
}

func part1(text string) int64 {
	transmission := hexadecimalToBinary(text)
	p, _ := parsePacket(transmission)
	return p.VersionSum()
}

func part2(text string) int64 {
	transmission := hexadecimalToBinary(text)
	p, _ := parsePacket(transmission)
	return p.Value()
}
