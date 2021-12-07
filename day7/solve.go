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

	fuel := []int{}
	for _, snp := range pos {
		np, _ := strconv.Atoi(snp)
		sum := 0
		for _, sp := range pos {
			p, _ := strconv.Atoi(sp)

			sum += abs(np - p)
		}
		fuel = append(fuel, sum)
		sum = 0
	}

	fmt.Println(min(fuel), fuel)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return x * (-1)
}

func min(x []int) int {
	m := math.MaxInt
	for _, i := range x {
		if i < m {
			m = i
		}
	}

	return m
}
