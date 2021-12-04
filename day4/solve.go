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
	bingo := strings.Split(content, "\n\n")
	bingo = bingo[:len(bingo)-1]

	numbers := strings.Split(bingo[0], ",")
	boards := parseBoards(bingo[1:])

	wt := []int{}
	for _, n := range numbers {
		for i, board := range boards {
			move(&board, n)
			if win(&board) {
				nmb, _ := strconv.Atoi(n)
				s := score(&board)
				if !already_win(wt, i) {
					wt = append(wt, i)
				}

				if len(wt) == len(boards) {
					fmt.Printf("board %d wins -> %d*%d=%d\n", i+1, nmb, s, nmb*s)
					return
				}
			}
		}
	}

}

func already_win(wt []int, w int) bool {
	for _, ws := range wt {
		if ws == w {
			return true
		}
	}
	return false
}

func score(board *[][]string) int {
	var s int

	for _, rows := range *board {
		for _, n := range rows {
			if n != "x" {
				i, _ := strconv.Atoi(n)
				s += i
			}
		}
	}

	return s
}

func move(board *[][]string, pos string) {
	for i, rows := range *board {
		for j, col := range rows {
			if col == pos {
				(*board)[i][j] = "x"
			}
		}
	}
}

func win(board *[][]string) bool {
	for i, row := range *board {
		if win_row(row) || win_col(*board, i) {
			return true
		}
	}

	return false
}

func win_row(row []string) bool {
	if "xxxxx" == strings.Join(row, "") {
		return true
	}

	return false
}

func win_col(cols [][]string, i int) bool {
	return cols[0][i] == "x" &&
		cols[1][i] == "x" &&
		cols[2][i] == "x" &&
		cols[3][i] == "x" &&
		cols[4][i] == "x"
}

func parseBoards(boards []string) [][][]string {
	a := [][][]string{}
	for _, board := range boards {
		b := [][]string{}
		for _, row := range strings.Split(board, "\n") {
			c := []string{}
			for _, d := range strings.Split(row, " ") {
				if d != "" && d != " " {
					c = append(c, d)
				}
			}
			b = append(b, c)
		}
		a = append(a, b)
	}

	return a
}
