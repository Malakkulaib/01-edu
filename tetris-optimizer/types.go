package main

type Point struct {
	X int
	Y int
}

type PuzzlePiece struct {
	Blocks []Point
	Letter rune
}
