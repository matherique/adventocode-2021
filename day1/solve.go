package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const in string = "in"

func main() {
	b, err := os.ReadFile(in)

	if err != nil {
		log.Fatal(err)
	}

	data := string(b)
	lines := strings.Split(data, "\n")

	size := len(lines)

	newlines := []string{}

	for i := range lines[:size-1] {
		if i+2 > size {
			continue
		}

		if lines[i] == "" || lines[i+1] == "" || lines[i+2] == "" {
			continue
		}

		a := toInt(lines[i])
		b := toInt(lines[i+1])
		c := toInt(lines[i+2])

		sum := strconv.Itoa(a + b + c)
		newlines = append(newlines, sum)
	}

	solve_2(newlines)

}

func solve_2(lines []string) {
	var (
		prev int
		inc  int
		dec  int
	)

	for i, v := range lines {
		if v == "" || i == 0 {
			continue
		}

		d := toInt(v)

		if d == prev {
			fmt.Printf("%d (no change)\n", d)
			prev = d
			continue
		}

		if d < prev {
			fmt.Printf("%d (decreassed)\n", d)
			dec += 1
		} else {
			inc += 1
			fmt.Printf("%d (increassed)\n", d)
		}

		prev = d
	}

	fmt.Println("inc: ", inc)
	fmt.Println("dec: ", dec)
}

func solve_1(lines []string) {
	var (
		prev int
		inc  int
		dec  int
	)

	for i, v := range lines {
		if v == "" || i == 0 {
			continue
		}

		d := toInt(v)

		if d < prev {
			fmt.Printf("%d (decreassed)", d)
			dec += 1
		} else {
			inc += 1
			fmt.Printf("%d (increassed)", d)
		}

		fmt.Println()

		prev = d
	}

	fmt.Println("inc: ", inc)
	fmt.Println("dec: ", dec)
}

func toInt(n string) int {
	d, err := strconv.Atoi(n)

	if err != nil {
		log.Fatalf("toInt(%s) error: %v\n", n, err)
	}

	return d
}
