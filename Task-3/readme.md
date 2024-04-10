# Golang Tasks

## Task 3 : Sudoku Solver

### Problem Statement

Given a filled 9x9 grid, your task is to check whether the given sudoku is solvable or not. The input for your program will be a 9x9 matrix where empty cells are represented by 0, and the filled cells are represented by their respective digits.

### Instructions 

1. Create a Golang program that takes the input Sudoku grid and returns whether the given sudoku is solvable or not.
2. If it fails then also return the row and column for which it fails
3. Test your solution with different Sudoku puzzles to validate its correctness.

### Solution

This Problem can be solved using maps, recursion and backtracking. Below are the steps that we can follow to find the solution of above problem.

> 1. Check if given sudoku is valid. If Yes then proceed to step 2 else return "Invalid Sudoku Provided" <br>
> 2. For Finding If Given Sudoku is solvable or not, we actually need to solve the sudoku. Check the value that can be fitted at the '0' spot by seeing Vertically, Horizontally and in it's respective grid . If there is no value to apply then we will return false else there is always a chance to find a right solved sudoku using given algorithm. Below is the Algorithm for same.
> 3. Iterate sudoku cells from top left to bottom right in row by row fashion.( for eg. 00, 01,..,08,10,11,....and till end of the sudoku cell which is 88 ) ( Here 00 describes row number 0 and column number 0 of sudoku matrix )
> 4. If Given Cell is not Empty ( Value is already been provided by puzzle maker ) THEN go on to check next cells without doing anything. Also if given cell is last then we got answer and hence return true.
> 5. If Given Cell is Empty THEN try all possible values from 1 to 9 inclusive. For Any Given Value, Check If we can insert it to sudoku -> If Yes then insert it and check if it produces a valid sudoku or not. -> If it produces valid sudoku then immediatly return true as we got our answer. ( we only want single solution ha ha). If it doesn't produces valid sudoku then backtrack from this value and try next value till 9 and if nothing works then return false.
> 6. Hence In this way we can check if given sudoku is solvable or not.

### Running Code
1. Clone the repo and Enter Task-3
2. cd Task-3
4. With the Go Setup in System , Command the terminal as : go run checkSudoku.go