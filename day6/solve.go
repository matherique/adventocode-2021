package main

import (
	"fmt"
	"log"
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
	lines := strings.Split(content[:len(content)-1], ",")

	pop := parse(lines)
	var day int
	limit := 80

	for {
		// fmt.Printf("%d day(s): %v\n", day, pop)
		if day == limit {
			break
		}

		for k, f := range pop {
			nv := f - 1
			if nv == -1 {
				nv = 6
			}

			if f == 0 {
				pop = append(pop, 8)
			}

			pop[k] = nv
		}

		day += 1
	}

	fmt.Println("total =", len(pop))
}

func parse(list []string) []int {
	res := []int{}
	for _, l := range list {
		n, _ := strconv.Atoi(l)
		res = append(res, n)
	}

	return res
}
