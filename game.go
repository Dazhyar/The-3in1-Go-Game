/* Explanation for the code used in this program that 
was not covered in class:

I spent my free time studying golang during the semester.
All the code related to time, math/rand, range keyword, append method,
slices, arrays, maps and the blank identifier I learned by myself and 
I did not use any external sources. The only thing that came from the
internet was the design for the hangman character that was done in 
ASCII art. I did not get the code of the game from the internet 
I only got inspiration for the design of hangman character.
*/

package main

import (
	"fmt"
	"time"
	"math/rand"
	"game/hangman"
	"game/tictactoe"
	"game/blackjack"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Creating a random seed based on the current time
	playerScore := 0

	printGameMessage()
	fmt.Println("Are you ready? You have 5 tries to guess the word correctly.\n")
	playerScore += hangman.Hangman()

	playerScore += tictactoe.TicTacToe()

	playerScore += blackjack.Blackjack()
	
	fmt.Printf("Your score is %d/3\n", playerScore)
	if playerScore == 3 {
		fmt.Println("Congratulations. You Won!")
		return 
	}

	fmt.Println("You Lost.")
}

/* This function goes through each byte in the gameMessage
   string and outputs the string version of it in 0.015 second
   intervals to make it look cool I guess. 
*/ 
func printGameMessage() {
	gameMessage := "Welcome to the almost impossible to beat game.\nIn this game you get to play three games.\n1. Hangman\n2. Tic-Tac-Toe\n3. Blackjack\nEach game is harder than the one before it.\nYou get one point from completing each game.\nYou need to win all three games to win.\nGood luck. You will need it!.\n\n"

	for i := range gameMessage { // i is set to be the starting byte index for each unicode character in the string
		fmt.Print(string(gameMessage[i]))
		time.Sleep(time.Millisecond * 15)
	}
}
