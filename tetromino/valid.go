package tetromino

// directions define the possible neighbours of a #.
var directions = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

// AreTetrominosValid checks if all tetrominos in the given slice are valid.
// It returns true if all tetrominos are valid; otherwise, false.
func AreTetrominosValid(tets []Tetromino) bool {
	for _, t := range tets {
		if !isValidTetromino(t) {
			return false
		}
	}
	return true
}

// isValidTetromino checks if a single Tetromino is valid according to
// the rules:
// - Each tetromino must be composed of exactly 4 '#' characters.
// - Each '#' character must be connected to at least one other '#' character.
// - A minimum of 6 connection.
// It returns true if the Tetromino is valid; otherwise, false.
func isValidTetromino(t Tetromino) bool {
	numConnections := 0
	numHashtags := 0

	for row, line := range t.Shape {
		for col, char := range line {

			if char != '#' && char != '.' {
				return false
			}

			if char == '#' {
				connectedHashtags := countConnectedHashtags(t.Shape, row, col)
				// Count total '#' and valid connections
				numHashtags++
				if connectedHashtags == 0 {
					return false
				} else {
					numConnections += connectedHashtags
				}
			}
		}
	}
	// return false if invalid
	if numConnections < 6 || numHashtags != 4 {
		return false
	}
	return true
}

// countConnectedHashtags counts the number of '#' characters connected in the 
// different directions defined in var directions.
func countConnectedHashtags(shape []string, row, col int) int {
	connectedHashtags := 0
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newRow < len(shape) && newCol >= 0 && newCol < len(shape[row]) {
			if shape[newRow][newCol] == '#' {
				connectedHashtags++
			}
		}
	}

	return connectedHashtags
}
