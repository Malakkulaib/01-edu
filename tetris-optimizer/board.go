package main

import (
	"fmt"
	"math"
)

func minSquareSize(n int) int {
	return int(math.Ceil(math.Sqrt(float64(n * 4))))
}

func makeBoard(size int) [][]rune {
	b := make([][]rune, size)

	for i := range b {
		b[i] = make([]rune, size)

		for j := range b[i] {
			b[i][j] = '.'
		}
	}

	return b
}

func printBoard(board [][]rune) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}
