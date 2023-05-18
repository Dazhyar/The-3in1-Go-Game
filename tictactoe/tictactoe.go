package tictactoe

import (
	"fmt"
	"math/rand"
)

func TicTacToe() int {
	board := makeBoard() // To get a new board at the beginning of the game.

	printBoard(board)
	for {
		selectCell(&board)

		// Here we will print what the board looks like after each round.
		printBoard(board)
	
		// The following switch statement is to check for the 8 possible winning combos.
		if areCellsEqual, winningCellIndex := checkForCellsEquality(board); areCellsEqual {
			return checkForTicTacToeWinner(board, winningCellIndex)
		}		
	}
}

// This function creates a new tic-tac-toe board using a string array of length 9.
func makeBoard() [9]string {
	board := [9]string{}
	for i := range board {
		board[i] = "[   ]"
	}

	return board
}

// This function prints out the tic-tac-toe board.
func printBoard(board [9]string) {
	for i := range board {
		fmt.Print(board[i])
		if (i+1) % 3 == 0 { // I check to see if the sum of 1 and the index of a cell from the board is divisble by 3 to print a new line every three cells.
			fmt.Println()
		}
	}
}

func selectCell(boardPtr *[9]string) {
	cellIndex := 0

	fmt.Print("Choose a cell from [1-9]:") // This for loop will give the player a cell
	for {
		fmt.Scanln(&cellIndex)
		if boardPtr[cellIndex-1] == "[   ]" {
			boardPtr[cellIndex-1] = "[ X ]"
			break
		}
		fmt.Print("Cell already occupied. Try another one: ")
	}
	checkForDraw(boardPtr)

	for { 
		cellIndex = rand.Intn(len(boardPtr)) // This will choose a random cell from the board for the computer
		if boardPtr[cellIndex] == "[   ]" { 
			boardPtr[cellIndex] = "[ O ]"
			break
		}
	}
	checkForDraw(boardPtr)
}

func checkForDraw(boardPtr *[9]string) {
	for i := range boardPtr {
		if boardPtr[i] == "[   ]" {
			return
		}
	}

	printBoard(*boardPtr)
	fmt.Println("There was a draw.")
	*boardPtr = makeBoard()
}

func checkForCellAvailabilty(chosenIndex int, board [9]string) bool {
		if board[chosenIndex] == "[  ]" {	
			return true
		}

		return false
}

func checkForCellsEquality(board [9]string) (bool, int) {
	winningCombos := [][]int{
		[]int{0, 1, 2}, []int{3, 4, 5}, 
		[]int{6, 7, 8}, []int{0, 3, 6},
		[]int{1, 4, 7}, []int{2, 5, 8},
		[]int{0, 4, 8}, []int{2, 4, 6},
	}

	for i := range winningCombos {
		if board[winningCombos[i][0]] == board[winningCombos[i][1]] && board[winningCombos[i][0]] == board[winningCombos[i][2]] && board[winningCombos[i][0]] != "[   ]" {
			return true, winningCombos[i][0]
		}
	}

	return false, 0
}

func checkForTicTacToeWinner(board [9]string, winningCellIndex int) int {
	if board[winningCellIndex] == "[ X ]" {
		fmt.Println("you won!\n")
		return 1
	}
	
	fmt.Println("You lost.\n")
	return 0
}
