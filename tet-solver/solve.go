package tetsolver

import (
	"math"

	"tetris-optimizer/tetromino"
)

// SolveSquare attempts to place terominos on the smallest square grid  possible,
// where all tetrominos  can be placed without overlapping.
// If that fails, it increases the square size and tries again until its succesful.
// It returns the solved square grid as a 2D slice of strings.
func SolveSquare(t []tetromino.Tetromino) [][]string {
	// Calculate the smallest possible square to try placement
	squareSize := int(math.Ceil(math.Sqrt(float64(len(t) * 4))))

	for {
		square := make([][]string, squareSize)
		for i := range square {
			square[i] = make([]string, squareSize)
			for j := range square[i] {
				square[i][j] = "."
			}
		}
		if placeTetrominos(square, t, 0) {
			return square
		}

		// If placement failed, increase the square size and try again
		squareSize++
	}
}
