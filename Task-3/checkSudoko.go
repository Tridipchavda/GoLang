package main

import "fmt"

// function to Check that sudoku is Valid and contain no '0' elements
func checkSudokuIsValidAndSolved(sudoku [9][9]int) (x int, y int, b bool) {

	// Iterating through matrix to check each element is at right place
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku); j++ {

			m := map[int]int{}
			// Checking vertical direction
			for k := 0; k < len(sudoku); k++ {
				m[sudoku[i][k]]++
			}
			for k := 0; k < len(sudoku); k++ {
				if m[sudoku[i][k]] != 1 {
					fmt.Println("By Vertical")
					return i, j, false
				}
			}

			m = map[int]int{}
			// Checking horizontal direction
			for k := 0; k < len(sudoku); k++ {
				m[sudoku[k][j]]++
			}
			for k := 0; k < len(sudoku); k++ {
				if m[sudoku[k][j]] != 1 {
					fmt.Println("By Horizontal")
					return i, j, false
				}
			}

			// Checking for 3X3 Grid from 9X9 Grid
			refI := 3 * (i / 3)
			refJ := 3 * (j / 3)

			m = map[int]int{}
			for k := refI; k < refI+3; k++ {
				for l := refJ; l < refJ+3; l++ {
					m[sudoku[k][l]]++
				}
			}
			for k := 0; k < len(sudoku); k++ {
				if m[sudoku[k][j]] != 1 {
					fmt.Println(sudoku[k][j])
					fmt.Println("By Grid")
					return i, j, false
				}
			}
		}
	}

	return -1, -1, true
}

// Function to solve the Sudoku problem with '0's in it
func checkSudokoSolvable(sudoku *[9][9]int) bool {

	// Iterating to check '0' in sudoku
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku); j++ {
			if sudoku[i][j] == 0 {
				m1 := map[int]int{}
				// Mapping vertical direction
				for k := 0; k < len(sudoku); k++ {
					m1[sudoku[i][k]]++
				}

				// Mapping horizontal direction
				m2 := map[int]int{}
				for k := 0; k < len(sudoku); k++ {
					m2[sudoku[k][j]]++
				}

				// Mapping for 3X3 Grid
				refI := 3 * (i / 3)
				refJ := 3 * (j / 3)

				m3 := map[int]int{}
				for k := refI; k < refI+3; k++ {
					for l := refJ; l < refJ+3; l++ {
						m3[sudoku[k][l]]++
					}
				}

				// All three Maps hold value that is filled in horizontal direction , Vertical direction and Grid 3x3
				// If Combining all three maps indexes gives 0 (0+0+0) then it is available to add in sudoku
				for k := 1; k <= len(sudoku); k++ {
					if m1[k]+m2[k]+m3[k] == 0 {
						// Putting value to find solution and backtrack if not possible
						sudoku[i][j] = k
						poss := checkSudokoSolvable(sudoku)
						if poss {
							return true
						} else {
							sudoku[i][j] = 0
						}
					}
				}
				// return false if there is not valid number to put in sudoku
				return false
			}
		}
	}
	// return true if solvable
	return true
}

// Function to represnt the sudoku in format
func printSudoko(sudoku [9][9]int) {
	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku); j++ {
			fmt.Print(sudoku[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	//  {5, 3, 4, 6, 7, 8, 9, 1, 2},
	// 	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	// 	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	// 	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	// 	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	// 	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	// 	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	// 	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	// 	{3, 4, 5, 2, 8, 6, 1, 7, 9},

	// Sudoku matrix
	sudoku := [9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	// Checking for elements to fill and check sudoku by solving
	printSudoko(sudoku)
	checkSudokoSolvable(&sudoku)
	printSudoko(sudoku)

	// function to check sudoku is Valid and if not return the (x,y) points and boolean value
	x, y, b := checkSudokuIsValidAndSolved(sudoku)

	fmt.Println(x, y, b)
}
