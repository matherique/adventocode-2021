package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const input = "in"

func main() {
	b, err := os.ReadFile(input)

	if err != nil {
		log.Fatal(err)
	}

	content := string(b)
	pos := strings.Split(content[:len(content)-1], ",")
	m := max(pos)

	min := math.MaxInt
	i := 0
	for {
		if i == m {
			break
		}
		sum := 0

		for _, sp := range pos {
			p, _ := strconv.Atoi(sp)
			diff := abs(i - p)
			a := z(diff)

			sum += a
			// fmt.Printf("Move from %d to %d: %d fuel\n", p, np, a)
		}
		// fmt.Println("==========", sum, "==========")
		if sum < min {
			min = sum
		}

		sum = 0
		i += 1
	}

	fmt.Println(min)
}

func array(size int) []int {
	l := []int{}
	var i int
	for {
		if i == size {
			break
		}

		l = append(l, i)
		i += 1
	}

	return l
}

func z(max int) int {
	i, sum := 0, 0

	for {
		sum += i
		i += 1

		if i > max {
			break
		}
	}

	return sum
}

func max(list []string) int {
	m := -1
	for _, item := range list {
		i, _ := strconv.Atoi(item)
		if i > m {
			m = i
		}
	}

	return m
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return x * (-1)
}
