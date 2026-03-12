package main

import (
	"fmt"
	"strings"
)

func loadPuzzle(data string) ([]PuzzlePiece, error) {
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
