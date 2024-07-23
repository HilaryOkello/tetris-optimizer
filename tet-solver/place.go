// Package tetsolver provides functions for solving the Tetris placement problem
// by recursively placing tetrominos on a square grid.
package tetsolver

import "tetris-optimizer/tetromino"


// Args struct holds the arguments necessary for placing and removing tetrominos.
type Args struct {
	square    [][]string // square is the grid/board
	tetromino tetromino.Tetromino //tetromino is the current tetromino
	x, y      int // x, y are the cordinates in the square we're placing/removing a tetromino
}

// placeTetrominos attempts to place each tetromino recursively on the square grid.
// It returns true if all tetrominos are successfully placed, false otherwise.
func placeTetrominos(square [][]string, tetrominos []tetromino.Tetromino, index int) bool {
	if index == len(tetrominos) {
		return true // Base case: all tetrominos are placed successfully
	}
	for x := 0; x < len(square); x++ {
		for y := 0; y < len(square[x]); y++ {
			args := Args{
				square:    square,
				tetromino: tetrominos[index],
				x:         x,
				y:         y,
			}
			if canPlaceTetromino(args) {
				placeTetromino(args)
				if placeTetrominos(square, tetrominos, index+1) {
					return true
				}
				removeTetromino(args)
			}
		}
	}
	return false // Could not place tetromino at current index
}

// canPlaceTetromino checks if a tetromino can be placed at the current position (args.x, args.y).
// It returns true if the tetromino can be placed, false otherwise.
func canPlaceTetromino(args Args) bool {
	for _, pos := range args.tetromino.Positions {
		newX, newY := args.x+pos[0], args.y+pos[1]
		if newX >= len(args.square) || newY >= len(args.square) || args.square[newX][newY] != "." {
			return false
		}
	}
	return true
}

// placeTetromino places the tetromino at the current position (args.x, args.y) on the square.
func placeTetromino(args Args) {
	for _, pos := range args.tetromino.Positions {
		newX, newY := args.x+pos[0], args.y+pos[1]
		args.square[newX][newY] = args.tetromino.Letter
	}
}

// placeTetromino removes the tetromino at the current position (args.x, args.y) on the square.
func removeTetromino(args Args) {
	for _, pos := range args.tetromino.Positions {
		newX, newY := args.x+pos[0], args.y+pos[1]
		args.square[newX][newY] = "."
	}
}
