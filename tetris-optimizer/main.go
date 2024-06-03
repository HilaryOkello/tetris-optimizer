package main

import (
	//	"bufio"

	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type Tetromino struct {
	Index int
	Shape [][]string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("You need to provide only one argument, a path to a textfile)")
	}
	filePath := os.Args[1]
	fileName := filepath.Base(os.Args[1])
	if filepath.Ext(filePath) != ".txt" {
		log.Fatalf("%s must be a textfile (i.e have a \".txt\" extension)", fileName)
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if len(fileContent) == 0 {
		log.Fatal("Empty file. No tetrominoes")
	}

	tetrominoes, mss := parseTetrominos(string(fileContent))
	myResult := make([][]string, mss)
	for i := range myResult {
		myResult[i] = make([]string, mss)
		for j := range myResult[i] {
			myResult[i][j] = "."
		}
	}
	placeTetrominoes(myResult, tetrominoes, 0)
	printTetrominoes(myResult)
}

func parseTetrominos(content string) ([]Tetromino, int) {
	sliceTets := strings.Split(string(content), "\n\n")
	var tetrominoes []Tetromino
	var countHash int
	for _, tet := range sliceTets {
		var tetLines [][]string
		lines := strings.Split(tet, "\n")
		for _, line := range lines {
			var linestrings []string
			for _, char := range line {
				switch char {
				case '.':
					linestrings = append(linestrings, string(char))
				case '#':
					linestrings = append(linestrings, string(char))
					countHash++
				default:
					log.Fatal("Invalid tetrominoes")

				}
			}
			tetLines = append(tetLines, linestrings)
		}
		tetromino := Tetromino{
			Index: len(tetrominoes),
			Shape: tetLines,
		}
		tetrominoes = append(tetrominoes, tetromino)
	}
	minSquareSize := math.Sqrt(float64(countHash))
	if minSquareSize-float64(int(minSquareSize)) > 0 {
		minSquareSize = math.Ceil(minSquareSize)
	}
	return tetrominoes, int(minSquareSize)
}

func placeTetrominoes(square [][]string, tetrominoes []Tetromino, index int) bool {
	if index == len(tetrominoes) {
		return true
	}
	for i := 0; i < len(square); i++ {
		for j := 0; j < len(square[i]); j++ {
			if canPlaceTetromino(square, tetrominoes[index], i, j) {
				placeTetromino(square, tetrominoes[index], i, j)
				if placeTetrominoes(square, tetrominoes, index+1) {
					return true
				}
				removeTetromino(square, tetrominoes[index], i, j)
			}
		}
	}
	return false
}

func canPlaceTetromino(square [][]string, tetromino Tetromino, x, y int) bool {
	size := len(square)
	for i := 0; i < len(tetromino.Shape); i++ {
		for j := 0; j < len(tetromino.Shape[i]); j++ {
			if tetromino.Shape[i][j] == "#" {
				if (x+i) >= size || (y+i) >= size || square[i+x][i+y] != "." {
					return false
				}
			}
		}
	}
	return true
}

func placeTetromino(square [][]string, tetromino Tetromino, x, y int) {
	for i := 0; i < len(tetromino.Shape); i++ {
		for j := 0; j < len(tetromino.Shape[i]); j++ {
			if tetromino.Shape[i][j] == "#" {
				square[i+x][i+y] = string(rune('A' + tetromino.Index))
			}
		}
	}
}

func removeTetromino(square [][]string, tetromino Tetromino, x, y int) {
	for i := 0; i < len(tetromino.Shape); i++ {
		for j := 0; j < len(tetromino.Shape[i]); j++ {
			if tetromino.Shape[i][j] == "#" {
				square[i+x][i+y] = "."
			}
		}
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
