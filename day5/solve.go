package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const input = "example"
const size = 10

func main() {
	b, err := os.ReadFile(input)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	lines = lines[:len(lines)-1]

	table := [size * size]int{}

	for _, line := range lines {
		x1, y1, x2, y2 := parse(line)

		if x1 != x2 && y1 != y2 {
			continue
		}

		table[idx(x1, y1)] += 1

		// fmt.Println(x1, y1, x2, y2)

		for {
			switch true {
			case x1 < x2:
				x1 += 1
			case x2 < x1:
				x1 -= 1
			case y1 < y2:
				y1 += 1
			case y2 < y1:
				y1 -= 1
			}

			// fmt.Println(x1, y1, x2, y2)
			table[idx(x1, y1)] += 1

			if x1 == x2 && y1 == y2 {
				break
			}
		}
	}

	var total int

	for _, v := range table {
		if v > 1 {
			total += 1
		}
	}

	fmt.Println("total =", total)

}

func idx(x, y int) int {
	return y + x*size
}

func parse(line string) (int, int, int, int) {
	pos := strings.Split(line, " -> ")

	pos1 := strings.Split(pos[0], ",")
	pos2 := strings.Split(pos[1], ",")

	sx1, sy1 := pos1[0], pos1[1]
	sx2, sy2 := pos2[0], pos2[1]

	return toInt(sx1), toInt(sy1), toInt(sx2), toInt(sy2)
}

func toInt(n string) int {
	i, err := strconv.Atoi(n)
	if err != nil {
		fmt.Printf("could not parse %q to int: %v\n", n, err)
	}

	return i
}
