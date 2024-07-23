// Package main is the entry point for the application.
// It reads tetrominos from a file and assemble them
// in order to create the smallest square possible.
package main

import (
	"fmt"
	"os"

	tetsolver "tetris-optimizer/tet-solver"
	"tetris-optimizer/tetromino"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <path to txt file> | cat -e")
		return
	}
	tetrominos, err := tetromino.ReadTetrominosFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !tetromino.AreTetrominosValid(tetrominos) {
		fmt.Println("ERROR")
		return
	}

	square := tetsolver.SolveSquare(tetrominos)
	tetsolver.PrintTetrominos(square)
}
