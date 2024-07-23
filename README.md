# Tetris-optimizer

tetris-optimizer is a program written in Go that receives only one argument, a path to a text file which will contain a list of tetrominoes, and assemble them in order to create the smallest square possible.

## Requirements

- Go 1.22.1 or later installed. You can download it from [golang.org](https://golang.org/dl/).

## Installation

```bash
git clone "https://learn.zone01kisumu.ke/git/hilaokello/tetris-optimizer.git"
cd tetris-optimizer
```

## Usage

To use the program, run it with a single argument which is the path to a text file containing tetrominos. For example:

```bash
go run . examples/sample1.txt | cat -e
```

Ad the output

```bash
/tetris-optimizer$ go run . examples/sample1.txt | cat -e
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
/tetris-optimizer$ 
```

## Valid Tetromino File
A valid `.txt` file will have tetrominos presented in the following configuration:
- The shape of a tetromino is represented by the `#` 
character and all other spaces by a `.`.
- Each tetromino consists of 4 lines of length 4 each.
- Tetrominos are seperated by single empty lines

### Example

```bash
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

....
.##.
.##.
....
```

## Implementation Details

The program consists of the following components:

- Main Program (main.go):

  - Entry point that reads the input file and coordinates the solution.
  - Validates input and outputs errors for invalid tetrominos or file format.
  - Uses the tetromino package to read and validate tetrominos.
  - Uses the tetsolver package to assemble tetrominos into the smallest square possible and print the solution.

- tet-solver Package:
  - Contains logic to assemble tetrominos into a square.
  - Implements algorithms to find the optimal arrangement of tetrominos.

- tetromino Package:
  - Handles reading tetrominos from a file.
  - Provides functions to validate tetrominos.

## Running Tests

The project includes unit tests to ensure correctness of tetromino validation and square assembly. Tests are located in respective test files within the tetromino and tetsolver packages.
Future Improvements

## Collaboration

To contribute to future improvements, please create a pull request.

## License

This project is licensed under the [MIT License](LICENSE)

