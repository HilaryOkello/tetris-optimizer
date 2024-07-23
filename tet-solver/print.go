package tetsolver

import "fmt"

//PrintTetrominos prints the square grid to the terminal. 
func PrintTetrominos(square [][]string) {
	for i := range square {
		for _, char := range square[i] {
			fmt.Printf("%s", char)
		}
		fmt.Println()
	}
}
