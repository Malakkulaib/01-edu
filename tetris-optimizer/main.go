package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

type PuzzlePiece struct {
	Blocks []Point
	Letter rune
}

func parseTetrominoes(data string) ([]PuzzlePiece, error) {
	data = strings.ReplaceAll(data, "\r", "")
	lines := strings.Split(strings.TrimSpace(data), "\n")

	var res []PuzzlePiece
	i := 0
	letter := 'A'

	for i < len(lines) {
		if i+4 > len(lines) {
			return nil, fmt.Errorf("bad format")
		}

		block := lines[i : i+4]
		points := []Point{}

		for y, line := range block {
			if len(line) != 4 {
				return nil, fmt.Errorf("bad line length")
			}
			for x, c := range line {
				if c != '.' && c != '#' {
					return nil, fmt.Errorf("bad char")
				}
				if c == '#' {
					points = append(points, Point{X: x, Y: y})
				}
			}
		}

		if len(points) != 4 || !isConnected(points) {
			return nil, fmt.Errorf("invalid tetromino")
		}

		points = normalize(points)
		res = append(res, PuzzlePiece{Blocks: points, Letter: letter})
		letter++
		i += 4

		if i < len(lines) && lines[i] == "" {
			i++
		}
	}

	return res, nil
}

func isConnected(points []Point) bool {
	links := 0
	for _, p := range points {
		for _, q := range points {
			if (abs(p.X-q.X) == 1 && p.Y == q.Y) ||
				(abs(p.Y-q.Y) == 1 && p.X == q.X) {
				links++
			}
		}
	}
	return links >= 6
}

func normalize(points []Point) []Point {
	minX, minY := points[0].X, points[0].Y
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
	}
	for i := range points {
		points[i].X -= minX
		points[i].Y -= minY
	}
	return points
}

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

func solve(board [][]rune, tetros []PuzzlePiece, idx int) bool {
	if idx == len(tetros) {
		return true
	}

	t := tetros[idx]
	size := len(board)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if canPlace(board, t, x, y) {
				place(board, t, x, y)
				if solve(board, tetros, idx+1) {
					return true
				}
				remove(board, t, x, y)
			}
		}
	}
	return false
}

func canPlace(board [][]rune, t PuzzlePiece, x, y int) bool {
	size := len(board)
	for _, p := range t.Blocks {
		nx, ny := x+p.X, y+p.Y
		if nx < 0 || ny < 0 || nx >= size || ny >= size {
			return false
		}
		if board[ny][nx] != '.' {
			return false
		}
	}
	return true
}

func place(board [][]rune, t PuzzlePiece, x, y int) {
	for _, p := range t.Blocks {
		board[y+p.Y][x+p.X] = t.Letter
	}
}

func remove(board [][]rune, t PuzzlePiece, x, y int) {
	for _, p := range t.Blocks {
		board[y+p.Y][x+p.X] = '.'
	}
}

func printBoard(board [][]rune) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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

	tetros, err := parseTetrominoes(string(content))
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
