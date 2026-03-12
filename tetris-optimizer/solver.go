package main

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

		nx := x + p.X
		ny := y + p.Y

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
