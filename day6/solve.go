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
	limit := 256

	// this solution is much better then the part 1
	for {
		if day == limit {
			break
		}

		n0, n1, n2, n3, n4, n5, n6, n7, n8 := pop[0], pop[1], pop[2], pop[3], pop[4], pop[5], pop[6], pop[7], pop[8]

		pop[0] = n1
		pop[1] = n2
		pop[2] = n3
		pop[3] = n4
		pop[4] = n5
		pop[5] = n6
		pop[6] = n7 + n0
		pop[7] = n8
		pop[8] = n0

		day += 1
	}

	fmt.Println(pop)
	var sum uint64
	for _, v := range pop {
		sum += v
	}

	fmt.Println("total =", sum)
}

func parse(list []string) [9]uint64 {
	res := [9]uint64{}

	for _, i := range list {
		n, _ := strconv.Atoi(i)
		res[n] += 1
	}

	return res
}
