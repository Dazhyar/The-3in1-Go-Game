package tictactoe

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"strconv"
)

func TicTacToe() int {
	scanner := bufio.NewScanner(os.Stdin)
	board := makeBoard() // To get a new board at the beginning of the game.

	printBoard(board)
	for {
		for {
			fmt.Print("Choose a box (1-9): " ) // Letting the player choose one of the cells
			scanner.Scan()
			
			playerIndex, _ := strconv.Atoi(scanner.Text())
			if board[playerIndex-1] == "[   ]" {
				board[playerIndex-1] = "[ X ]"
				break
			}

			fmt.Println("Try another cell.")
		}

		allBoardUsed := true
		for i := range board {
			if board[i] == "[   ]" {
				allBoardUsed = false
				fmt.Println("he")
				break
			}
		}

		if allBoardUsed {
			fmt.Println("hi")
			board = makeBoard() // We will create a new empty board if nobody won. 
		}

		for {
			randIndex := rand.Intn(9) // This will choose a random cell from the board for the computer
			if board[randIndex] == "[   ]" { // If that randomly picked cell is not occupied then we will put an O in it as the computer's choice. 
				board[randIndex] = "[ O ]"
				break
			}
		}
		printBoard(board) // Here we will print what the board looks like after each round.

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

// This function prints out the tic-tac-toe form.
func printBoard(board [9]string) {
	for i := range board {
		fmt.Print(board[i])
		if (i+1) % 3 == 0 { // I check to see if the sum of 1 and the index of a cell from the board is divisble by 3 to print a new line every three cells.
			fmt.Println()
		}
	}
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
