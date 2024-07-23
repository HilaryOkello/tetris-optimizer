// Package tetromino provides functionality to read, process, and check validity of
// tetromino configurations from a text file. 
// Tetrominos are represented as shapes composed of '#'.
package tetromino

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Tetromino represents properties of a tetromino.
type Tetromino struct {
	Letter    string   // Letter is the identifier of the tetromino (e.g., 'A', 'B', etc.).
	Shape     []string // Shape holds the representation of the tetromino as in the file.
	Positions [][2]int //Positions stores the coordinates of the '#' offset to top-most, left-most pos.
}

// ReadTetrominosFile reads tetrominos from a text file and returns a slice of Tetromino structs.
// The file path is expected as the first argument in os.Args.
func ReadTetrominosFile() ([]Tetromino, error) {
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
	var tetrominos []Tetromino
	var tetLines []string
	terominoLines := 4 // Expected number of lines per tetromino shape.
	var emptyLines int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			emptyLines++
			continue
		}

		if emptyLines > 1 {
			return nil, fmt.Errorf("ERROR")
		}

		// Each line in the tetromino shape must be exactly 4 characters.
		if len(line) != 4 {
			return nil, fmt.Errorf("ERROR")
		}

		emptyLines = 0

		tetLines = append(tetLines, line)

		if len(tetLines) == terominoLines {
			tetromino := createTetromino(tetLines, len(tetrominos))
			tetrominos = append(tetrominos, tetromino)
			tetLines = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ERROR")
	}

	if len(tetLines) > 0 {
		fmt.Println(len(tetLines))
		return nil, fmt.Errorf("ERROR")
	}

	return tetrominos, nil
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
		Letter:    string('A' + index),
		Shape:     tetLines,
		Positions: positions,
	}
}
