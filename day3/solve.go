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
	f, err := os.ReadFile(input)

	if err != nil {
		log.Fatalln(err)
	}

	content := string(f)
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1]

	sum := make([]int, len(lines[0]))

	for _, line := range lines {
		for i, n := range strings.Split(line, "") {
			if n == " " {
				continue
			}
			nv, _ := strconv.Atoi(n)
			sum[i] += nv
		}
	}

	size := len(lines)
	var most, least string

	for _, i := range sum {
		if i < (size / 2) {
			most += "0"
			least += "1"
		} else {
			most += "1"
			least += "0"
		}
	}

	m, _ := strconv.ParseInt(most, 2, 64)
	l, _ := strconv.ParseInt(least, 2, 64)

	fmt.Println(m * l)
}

func commoms(list []string) (string, string) {
	var ones, zeros int

	for _, i := range list {
		if i == "1" {
			ones += 1
			continue
		}

		zeros += 1
	}

	if ones > zeros {
		return "1", "0"
	}

	return "0", "1"
}
