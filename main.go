package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"
)

type Tetromino struct {
	Index     int
	Shape     []string
	MinRow    int
	MinCol    int
	Positions [][2]int
}

func main() {
	t := time.Now()
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
	square := MakeSquare(tetrominoes)
	printTetrominoes(square)
	fmt.Println(time.Since(t))
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
				tetromino := createTetromino(tetLines, len(tetrominoes))
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

func createTetromino(tetLines []string, index int) Tetromino {
	minRow, minCol := 3, 3
	var positions [][2]int

	for row, line := range tetLines {
		for col, char := range line {
			if char == '#' {
				if row < minRow {
					minRow = row
				}
				if col < minCol {
					minCol = col
				}
				positions = append(positions, [2]int{row, col})
			}
		}
	}

	// Adjust positions relative to minRow and minCol
	for i := range positions {
		positions[i][0] -= minRow
		positions[i][1] -= minCol
	}

	return Tetromino{
		Index:     index,
		Shape:     tetLines,
		MinRow:    minRow,
		MinCol:    minCol,
		Positions: positions,
	}
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
		for y := 0; y < len(square); y++ {
			if canPlaceTetromino(square, tetrominoes[index], x, y) {
				placeTetromino(square, tetrominoes[index], x, y, index)
				if placeTetrominoes(square, tetrominoes, index+1) {
					return true
				}
				removeTetromino(square, tetrominoes[index], x, y)
			}
		}
	}
	return false
}

func canPlaceTetromino(square [][]string, tetromino Tetromino, x, y int) bool {
	for _, pos := range tetromino.Positions {
		newX, newY := x+pos[0], y+pos[1]
		if newX >= len(square) || newY >= len(square) || square[newX][newY] != "." {
			return false
		}
	}
	return true
}

func placeTetromino(square [][]string, tetromino Tetromino, x, y, index int) {
	for _, pos := range tetromino.Positions {
		newX, newY := x+pos[0], y+pos[1]
		square[newX][newY] = string(rune('A' + index))
	}
}

func removeTetromino(square [][]string, tetromino Tetromino, x, y int) {
	for _, pos := range tetromino.Positions {
		newX, newY := x+pos[0], y+pos[1]
		square[newX][newY] = "."
	}
}

func printTetrominoes(square [][]string) {
	for i := range square {
		for _, char := range square[i] {
			fmt.Printf("%s", char)
		}
		println()
	}
}
