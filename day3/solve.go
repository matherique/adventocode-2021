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

	lines_most := lines
	lines_least := lines

	size := len(lines[0]) - 1
	var i int

	// calculate oxygen generator rating
	for {
		most, _ := mostleast(lines_most)

		m := string(most[i])

		ml := []string{}

		for _, line := range lines_most {
			if string(line[i]) == m {
				ml = append(ml, line)
			}
		}

		lines_most = ml

		i += 1

		if i > size || len(lines_most) == 1 {
			break
		}
	}

	i = 0
	// calculate CO2 scrubber rating
	for {
		_, least := mostleast(lines_least)

		m := string(least[i])

		ll := []string{}

		for _, line := range lines_least {
			if string(line[i]) == m {
				ll = append(ll, line)
			}
		}

		lines_least = ll

		i += 1

		if i > size || len(lines_least) == 1 {
			break
		}
	}

	ogr := lines_most[0]
	csr := lines_least[0]

	ogrI, _ := strconv.ParseInt(ogr, 2, 64)
	csrI, _ := strconv.ParseInt(csr, 2, 64)

	fmt.Println(ogrI, csrI)

	fmt.Println(ogrI * csrI)
}

func remove(s []string, idx int) []string {
	return append(s[:idx], s[idx+1:]...)
}

func mostleast(lines []string) (string, string) {
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
	var most string
	var least string

	for _, i := range sum {
		switch true {
		case size-i > i:
			least += "1"
			most += "0"
		case size-i < i:
			least += "0"
			most += "1"
		case size-i == i:
			least += "0"
			most += "1"
		}
	}

	return most, least
}
