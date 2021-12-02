package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const file string = "in"

func main() {
	file, err := os.ReadFile(file)

	if err != nil {
		log.Fatalln(err)
	}

	content := string(file)
	lines := strings.Split(content, "\n")

	var (
		d int
		h int
	)

	for _, line := range lines[:len(lines)-1] {
		l := strings.Split(line, " ")
		m := l[0]
		n, _ := strconv.Atoi(l[1])

		switch m {
		case "forward":
			h += n
		case "up":
			d -= n
		case "down":
			d += n
		}
	}

	fmt.Println("d ", d)
	fmt.Println("h ", h)
	fmt.Println("dxh=", d*h)
}
