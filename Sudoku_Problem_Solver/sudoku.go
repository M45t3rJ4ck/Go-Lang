package main

import (
	"fmt"
	"strconv"
)

// 2) Create a function to display board:
func displayBoard(puzz [][]int) {
	// .1) Get length of columns.
	puzz_cols := len(puzz[0])

	// .2) Cycle through rows
	for i := 0; i < len(puzz); i++ {
		// .2.1) Cycle through columns
		for j := 0; j < puzz_cols; j++ {
			// .2.1.1) Print Board.
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		fmt.Println()
	}
}

// 4) Find all 0's on the board and return them as empty spaces:
func getEmptySpace(puzz [][]int) (int, int) {
	// .1) Get length of columns.
	puzz_cols := len(puzz[0])

	// .2) Cycle through rows
	for i := 0; i < len(puzz); i++ {
		// .2.1) Cycle through columns
		for j := 0; j < puzz_cols; j++ {
			// .2.1.1) Find the 0's
			if puzz[i][j] == 0 {
				// .2.1.1.1) Return that space.
				return i, j
			}
		}
	}
	// .3) Puzz completed and no 0's left.
	return -1, -1
}

// 6) Check if rows contain numbers 1 - 9:
func isNumValid(puzz [][]int, guess int, row int, cols int) bool {
	// .1) Find the index
	for index := range puzz {
		// .1.1) Check if row at index is not a guess
		if puzz[row][index] == guess && cols != index {
			return false
		}
	}
	// .3) Go through range of puzz again
	for index := range puzz {
		// .1) Check if column at index is not a guess.
		if puzz[index][cols] == guess && cols != index {
			return false
		}
	}
	// .4) Is number valid for box
	// .4.1) Cycle through rows 1 to 3 with -> row - (row % 3) + value for cycling = 0 - (0 % 3) + 0
	for k := 0; k < 3; k++ {
		// .1.2) Cycle through cols 1 to 3 with -> cols - (cols % 3) + value for cycling = 0 - (0 % 3) + 0
		for l := 0; l < 3; l++ {
			if puzz[row-row%3+k][cols-cols%3+l] == guess && (row-row%3+k != row || cols-cols%3+l != cols) {
				return false
			}
		}
	}

	// .2) Otherwise number was not found anywhere else.
	return true
}

// 8) Implement Recursion to solve sudoku
func solvePuzzle(puzz [][]int) bool {
	row, cols := getEmptySpace(puzz)

	if row == -1 {
		// .1) Puzzle solved
		return true
	} else {
		row, cols := getEmptySpace(puzz)
	}

	for i := 1; i <= 9; i++ {
		if isNumValid(puzz, i, row, cols) {
			puzz[row][cols] = i

			if solvePuzzle(puzz) {
				return true
			}
			puzz[row][cols] = 0
		}
	}
	return false
}

func main() {
	// 1) Create Multi-Dimensional Slice of board:
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	// 3) Test if board displays:
	displayBoard(puzz)

	// 5) Test if empty space can be found:
	// .5.1) Get empty spaces.
	row, _ := getEmptySpace(puzz)
	// .5.2) Cycle till first empty space
	if row != -1 {
		//.5.2.1) Print first empty space
		fmt.Println(getEmptySpace(puzz))
	} else {
		// .5.2.2) Puzzle solved.
		fmt.Println("Puzzle is solved!")
	}

	// 7) Test the numbers per row:
	// .7.1) Check the rows.
	fmt.Println(isNumValid(puzz, 1, 0, 0))
	fmt.Println(isNumValid(puzz, 7, 0, 0))
	//.7.2) Check the columns.
	fmt.Println(isNumValid(puzz, 6, 0, 0))
	fmt.Println(isNumValid(puzz, 9, 0, 0))
	//.7.3) Check the box.
	fmt.Println(isNumValid(puzz, 7, 4, 0))
	fmt.Println(isNumValid(puzz, 1, 0, 6))

	// 9) Test Solve Puzzle
	displayBoard(puzz)
	fmt.Println()
	solvePuzzle(puzz)
	displayBoard(puzz)
}
