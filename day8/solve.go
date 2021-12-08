package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const input = "in"

var (
	ds = []int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
)

func main() {
	f, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var sum int
	for s.Scan() {
		line := strings.Split(s.Text(), " | ")
		// us := strings.Split(line[0], " ")
		fd := strings.Split(line[1], " ")

		for _, f := range fd {
			if isUniq(f) {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}

func isUniq(f string) bool {
	return len(f) == ds[1] || len(f) == ds[4] || len(f) == ds[7] || len(f) == ds[8]
}
