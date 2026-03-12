package main

import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	tetros, err := loadPuzzle(string(content))
	if err != nil || len(tetros) == 0 {
		fmt.Println("ERROR")
		return
	}

	size := minSquareSize(len(tetros))
	for {
		board := makeBoard(size)
		if solve(board, tetros, 0) {
			printBoard(board)
			return
		}
		size++
	}
}
