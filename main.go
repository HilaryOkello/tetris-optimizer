package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
)

type Tetromino struct {
	Index int
	Shape []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . sample.txt | cat -e")
		return
	}
	tetrominoes, err := ReadTetrominoesFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	if !AreTetrominoesValid(tetrominoes) {
		fmt.Println("ERROR")
		return
	}
	fmt.Println(tetrominoes)
	square := MakeSquare(tetrominoes)
	printTetrominoes(square)
}

func ReadTetrominoesFile() ([]Tetromino, error) {
	filePath := os.Args[1]
	if filepath.Ext(filePath) != ".txt" {
		return nil, fmt.Errorf("ERROR")
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ERROR")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tetrominoes []Tetromino
	tetLines := []string{}
	var count int
	var empty int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 4 {
			if empty > 1 {
				return nil, fmt.Errorf("ERROR")
			} else {
				empty = 0
			}
			count++
			tetLines = append(tetLines, line)
			if count == 4 {
				tetromino := Tetromino{
					Index: len(tetrominoes),
					Shape: tetLines,
				}
				tetrominoes = append(tetrominoes, tetromino)
				tetLines = []string{}
				count = 0
			}
		}
		if line == "" {
			empty++
		}
	}
	if count != 0 {
		return nil, fmt.Errorf("ERROR")
	}
	return tetrominoes, nil
}

func AreTetrominoesValid(tets []Tetromino) bool {
	for _, t := range tets {
		countConnections := 0
		countHashtags := 0

		for row, line := range t.Shape {
			for col, char := range line {
				connectedHashtags := 0

				if char != '#' && char != '.' {
					return false
				}
				if char == '#' {
					// Define if above, below,left, and right == # & count connections
					above := row > 0 && t.Shape[row-1][col] == '#'
					below := row < len(t.Shape)-1 && t.Shape[row+1][col] == '#'
					left := col > 0 && t.Shape[row][col-1] == '#'
					right := col < len(line)-1 && t.Shape[row][col+1] == '#'
					// Count #s
					countHashtags++
					if above {
						connectedHashtags++
					}
					if below {
						connectedHashtags++
					}
					if left {
						connectedHashtags++
					}
					if right {
						connectedHashtags++
					}
					if connectedHashtags == 0 {
						return false
					} else {
						countConnections += connectedHashtags
					}
				}
			}
		}
		// return false if invalid
		if countConnections < 6 || countHashtags != 4 {
			return false
		}

	}
	return true
}

func MakeSquare(t []Tetromino) [][]string {
	// Start with the minimum possible size
	squareSize := int(math.Ceil(math.Sqrt(float64(len(t) * 4))))

	for {
		// Create a square of the current size
		square := make([][]string, squareSize)
		for i := range square {
			square[i] = make([]string, squareSize)
			for j := range square[i] {
				square[i][j] = "."
			}
		}

		// Try to place all tetrominoes
		if placeTetrominoes(square, t, 0) {
			return square
		}

		// If placement failed, increase the square size and try again
		squareSize++
	}
}

func placeTetrominoes(square [][]string, tetrominoes []Tetromino, index int) bool {
	if index == len(tetrominoes) {
		return true
	}
	for x := 0; x < len(square); x++ {
		for y := 0; y < len(square[x]); y++ {
			fmt.Printf("Trying to place Tet %d\n", index)
			if canPlaceTetromino(square, tetrominoes[index], x, y) {
				placeTetromino(square, tetrominoes[index], x, y)
				if placeTetrominoes(square, tetrominoes, index+1) {
					fmt.Printf("Successfuly placed Tet %d\n", index)
					return true
				}
				removeTetromino(square, tetrominoes[index], x, y)
			}
		}
	}
	fmt.Printf("Failed to place Tet %d\n", index)
	return false
}

func canPlaceTetromino(square [][]string, tetromino Tetromino, x, y int) bool {
	minRow, minCol := getTetrominoOffset(tetromino)
	for row, line := range tetromino.Shape {
		for col, char := range line {
			if char == '#' {
				newRow, newCol := x+(row-minRow), y+(col-minCol)
				if newRow >= len(square) || newCol >= len(square) || square[newRow][newCol] != "." {
					return false
				}
			}
		}
	}
	return true
}

func placeTetromino(square [][]string, tetromino Tetromino, x, y int) {
	minRow, minCol := getTetrominoOffset(tetromino)
	for row, line := range tetromino.Shape {
		for col, char := range line {
			if char == '#' {
				newRow, newCol := x+(row-minRow), y+(col-minCol)
				square[newRow][newCol] = string(rune('A' + tetromino.Index))
			}
		}
	}
}

func removeTetromino(square [][]string, tetromino Tetromino, x, y int) {
	minRow, minCol := getTetrominoOffset(tetromino)
	for row, line := range tetromino.Shape {
		for col, char := range line {
			if char == '#' {
				newRow, newCol := x+(row-minRow), y+(col-minCol)
				square[newRow][newCol] = "."
			}
		}
	}
}

func getTetrominoOffset(tetromino Tetromino) (int, int) {
	minRow, minCol := 3, 3
	for row, line := range tetromino.Shape {
		for col, char := range line {
			if char == '#' {
				if row < minRow {
					minRow = row
				}
				if col < minCol {
					minCol = col
				}
			}
		}
	}
	return minRow, minCol
}

func printTetrominoes(square [][]string) {
	for i := range square {
		for _, char := range square[i] {
			fmt.Printf("%s", char)
		}
		println()
	}
}
