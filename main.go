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
	square := MakeSquare(tetrominoes)
	placeTetrominoes(square, tetrominoes)
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
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 4 {
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
		} else if line == "" {
			if count != 0 {
				return nil, fmt.Errorf("ERROR")
			}
		} else {
			return nil, fmt.Errorf("ERROR")
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
	squareSize := math.Ceil(math.Sqrt(float64(len(t)) * 4))
	lineLength := squareSize
	square := [][]string{}
	line := []string{}
	for lineLength != 0 {
		line = append(line, ".")
		lineLength--
	}
	for squareSize != 0 {
		square = append(square, line)
		squareSize--
	}
	return square
}

func placeTetrominoes(square [][]string, tetrominoes []Tetromino) {
	index := 0
	for x := range square {
		for y := range square[x] {
			if canPlaceTetromino(square, tetrominoes[index], x, y) {
				placeTetromino(square, tetrominoes[index], x, y)
				index++
			}
		}
	}
}

func canPlaceTetromino(square [][]string, tetromino Tetromino, x, y int) bool {
	for row, line := range tetromino.Shape {
		for col, char := range line {
			if char == '#' {
				above := row > 0 && tetromino.Shape[row-1][col] == '#'
				below := row < len(tetromino.Shape)-1 && tetromino.Shape[row+1][col] == '#'
				left := col > 0 && tetromino.Shape[row][col-1] == '#'
				right := col < len(line)-1 && tetromino.Shape[row][col+1] == '#'
				if square[x][y] != "." {
					return false
				} else if above && x > 0 && square[x-1][y] != "." {
					return false
				} else if below && x < len(square)-1 && square[x+1][y] != "." {
					return false
				} else if left && y > 0 && square[x][y-1] != "." {
					return false
				} else if right && y < len(square[x])-1 && square[x][y+1] != "." {
					return false
				}

			}
		}
	}
	return true
}

func placeTetromino(square [][]string, tetromino Tetromino, x, y int) {
	for row, line := range tetromino.Shape {
		for col, char := range line {
			if char == '#' {
				square[x][y] = string(rune('A' + tetromino.Index))
				right := col < len(line)-1 && tetromino.Shape[row][col+1] == '#'
				if right && y < len(square[x])-1 && square[x][y+1] == "." {
					y += 1
				}
			}
		}
		x++
	}
}

// func removeTetromino(square [][]string, tetromino Tetromino, x, y int) {

// }

func printTetrominoes(square [][]string) {
	for i := range square {
		for _, char := range square[i] {
			fmt.Printf("%s", char)
		}
		println()
	}
}
