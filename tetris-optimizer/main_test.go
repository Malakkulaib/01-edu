package main

import (
	"strings"
	"testing"
)

// ============================================================
// BAD TESTS - loadPuzzle should return an error
// ============================================================

// bad00 - piece has 5 blocks instead of 4
func TestBad00_FiveBlocks(t *testing.T) {
	input := `####
...#
....
....`
	_, err := loadPuzzle(input)
	if err == nil {
		t.Error("expected error: piece has 5 blocks")
	}
}

// bad01 - diagonal blocks, not connected
func TestBad01_Diagonal(t *testing.T) {
	input := `...#
..#.
.#..
#...`
	_, err := loadPuzzle(input)
	if err == nil {
		t.Error("expected error: diagonal blocks are not connected")
	}
}

// bad02 - two blocks on top, two on bottom, not connected
func TestBad02_TwoAndTwo(t *testing.T) {
	input := `...#
...#
#...
#...`
	_, err := loadPuzzle(input)
	if err == nil {
		t.Error("expected error: blocks are not connected")
	}
}

// bad03 - empty piece, no blocks at all
func TestBad03_Empty(t *testing.T) {
	input := `....
....
....
....`
	_, err := loadPuzzle(input)
	if err == nil {
		t.Error("expected error: piece has 0 blocks")
	}
}

// bad04 - two blocks on top right, two on bottom left, not connected
func TestBad04_NotConnected(t *testing.T) {
	input := `..##
....
....
##..`
	_, err := loadPuzzle(input)
	if err == nil {
		t.Error("expected error: blocks are not connected")
	}
}

// bad05 - first two pieces are valid, third piece is not connected
func TestBad05_ThirdPieceBad(t *testing.T) {
	input := `...#
...#
...#
...#

....
....
....
####

#..#
....
....
....`
	_, err := loadPuzzle(input)
	if err == nil {
		t.Error("expected error: third piece is not connected")
	}
}

// ============================================================
// GOOD TESTS - loadPuzzle should succeed and solve should find a solution
// ============================================================

// good00 - single 2x2 square piece
func TestGood00_SingleSquare(t *testing.T) {
	input := `....
.##.
.##.
....`

	pieces, err := loadPuzzle(input)
	if err != nil {
		t.Fatalf("loadPuzzle returned error: %v", err)
	}
	if len(pieces) != 1 {
		t.Fatalf("expected 1 piece, got %d", len(pieces))
	}

	size := minSquareSize(len(pieces))
	board := makeBoard(size)
	if !solve(board, pieces, 0) {
		t.Fatal("no solution found")
	}

	if !validateBoard(board, pieces) {
		t.Error("board is invalid")
		printTestBoard(t, board)
	}
}

// good01 - three pieces
func TestGood01_ThreePieces(t *testing.T) {
	input := `...#
...#
...#
...#

....
....
....
####

.###
...#
....
....`

	pieces, err := loadPuzzle(input)
	if err != nil {
		t.Fatalf("loadPuzzle returned error: %v", err)
	}
	if len(pieces) != 3 {
		t.Fatalf("expected 3 pieces, got %d", len(pieces))
	}

	size := minSquareSize(len(pieces))
	for {
		board := makeBoard(size)
		if solve(board, pieces, 0) {
			if !validateBoard(board, pieces) {
				t.Error("board is invalid")
				printTestBoard(t, board)
			}
			return
		}
		size++
		if size > 20 {
			t.Fatal("no solution found, board size reached 20")
		}
	}
}

// good02 - four pieces
func TestGood02_FourPieces(t *testing.T) {
	input := `...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
.##.
.##.
....`

	pieces, err := loadPuzzle(input)
	if err != nil {
		t.Fatalf("loadPuzzle returned error: %v", err)
	}
	if len(pieces) != 4 {
		t.Fatalf("expected 4 pieces, got %d", len(pieces))
	}

	size := minSquareSize(len(pieces))
	for {
		board := makeBoard(size)
		if solve(board, pieces, 0) {
			if !validateBoard(board, pieces) {
				t.Error("board is invalid")
				printTestBoard(t, board)
			}
			return
		}
		size++
		if size > 20 {
			t.Fatal("no solution found, board size reached 20")
		}
	}
}

// good03 - twelve pieces
func TestGood03_TwelvePieces(t *testing.T) {
	input := `....
.##.
.##.
....

...#
...#
...#
...#

....
..##
.##.
....

....
.##.
.##.
....

....
..#.
.##.
.#..

.###
...#
....
....

##..
.#..
.#..
....

....
.##.
.##.
....

....
..##
.##.
....

##..
.#..
.#..
....

.#..
.##.
..#.
....

....
###.
.#..
....`

	pieces, err := loadPuzzle(input)
	if err != nil {
		t.Fatalf("loadPuzzle returned error: %v", err)
	}
	if len(pieces) != 12 {
		t.Fatalf("expected 12 pieces, got %d", len(pieces))
	}

	size := minSquareSize(len(pieces))
	for {
		board := makeBoard(size)
		if solve(board, pieces, 0) {
			if !validateBoard(board, pieces) {
				t.Error("board is invalid")
				printTestBoard(t, board)
			}
			return
		}
		size++
		if size > 20 {
			t.Fatal("no solution found, board size reached 20")
		}
	}
}

// hard - twelve pieces with a T-piece included
func TestHard_TwelvePieces(t *testing.T) {
	input := `....
.##.
.##.
....

.#..
.##.
.#..
....

....
..##
.##.
....

....
.##.
.##.
....

....
..#.
.##.
.#..

.###
...#
....
....

##..
.#..
.#..
....

....
.##.
.##.
....

....
..##
.##.
....

##..
.#..
.#..
....

.#..
.##.
..#.
....

....
###.
.#..
....`

	pieces, err := loadPuzzle(input)
	if err != nil {
		t.Fatalf("loadPuzzle returned error: %v", err)
	}
	if len(pieces) != 12 {
		t.Fatalf("expected 12 pieces, got %d", len(pieces))
	}

	size := minSquareSize(len(pieces))
	for {
		board := makeBoard(size)
		if solve(board, pieces, 0) {
			if !validateBoard(board, pieces) {
				t.Error("board is invalid")
				printTestBoard(t, board)
			}
			return
		}
		size++
		if size > 20 {
			t.Fatal("no solution found, board size reached 20")
		}
	}
}

// ============================================================
// Helper Functions
// ============================================================

// validateBoard checks that every piece appears exactly 4 times
// and that no unknown character exists on the board
func validateBoard(board [][]rune, pieces []PuzzlePiece) bool {
	size := len(board)

	for _, piece := range pieces {
		count := 0
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				if board[y][x] == piece.Letter {
					count++
				}
			}
		}
		if count != 4 {
			return false
		}
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if board[y][x] != '.' {
				found := false
				for _, piece := range pieces {
					if board[y][x] == piece.Letter {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
		}
	}

	return true
}

// printTestBoard prints the board to the test log for debugging
func printTestBoard(t *testing.T, board [][]rune) {
	t.Helper()
	var sb strings.Builder
	sb.WriteString("\nBoard:\n")
	for _, row := range board {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	t.Log(sb.String())
}
