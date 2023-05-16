package hangman

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
)

func Hangman() int {
	scanner := bufio.NewScanner(os.Stdin)
	words := []string{"cat", "city", "world", "table", "water"}
	secretWord, userGuess := words[rand.Intn(len(words))], ""
	
	for i := 0; i < len(secretWord); i++ {
		userGuess += "-"
	}
		
	tries := 5
	for tries > 0 {
		guessed := false

		printHangman(tries)
		fmt.Printf("The secret word: %s.\n", userGuess)
		fmt.Print("Guess one of the word's letters: ")
		scanner.Scan()

		for i := range secretWord { // This goes through the byte index for each character in the secretWord to check if the user's input is correct and is part of the secret word.
			if scanner.Text() == string(secretWord[i]) {
				userGuess = userGuess[:i] + string(secretWord[i]) + userGuess[i+1:] 
				guessed = true
			}			
		}
		fmt.Println()

		if !guessed { // After each wrong guess the player loses a try
			tries -= 1
			continue
		}

		if userGuess == secretWord {
			fmt.Printf("You guessed the word correctly! The word: %s.\n\n", secretWord)
			return 1
		}
	}

	fmt.Println("You lost! your hangman is ehh... well not alive exactly.\n")
	return 0
}

func printHangman(tries int) {
	fmt.Println("______________    ")
	fmt.Println("|           |     ")
	fmt.Println("|           |     ")
	printHead(tries)
	printHands(tries)
	printLegs(tries)
	fmt.Println("|                 ")
	fmt.Println("|_________________")
}

func printHead(tries int) {
	if tries > 0 {
		fmt.Println("|         ..... ")
		fmt.Println("|         .* *. ")
		fmt.Println("|         . - . ")
		fmt.Println("|         ..... ")
	} else {
		fmt.Println("|               ")
		fmt.Println("|               ")
		fmt.Println("|               ")
		fmt.Println("|               ")
	}
}

func printHands(tries int) {
	if tries > 2 {
		fmt.Println("|          /|\\ ")
		fmt.Println("|           |   ")
		fmt.Println("|           |   ")
	} else if tries > 1 {
		fmt.Println("|          /|   ")
		fmt.Println("|           |   ")
		fmt.Println("|           |   ")
	} else {
		fmt.Println("|               ")
		fmt.Println("|               ")
		fmt.Println("|               ")
	}
}

func printLegs(tries int) {
	if tries > 4 {
		fmt.Println("|          / \\")
	} else if tries > 3 {
		fmt.Println("|          /   ")
	} else {
		fmt.Println("|              ")
	}
}
