package blackjack

import (
	"fmt"	
	"math/rand"
	"strconv"
)

func Blackjack() int {
	deck :=  shuffleDeck()
	playerCards, computerCards := map[string]int{}, map[string]int{}
	playerPoints, computerPoints := 0, 0

	for i := 0; i < 3; i++ {
		randCardIndex := rand.Intn(len(deck))
		playerCards[deck[randCardIndex]], _ = strconv.Atoi(string(deck[randCardIndex][0])) 
		deck = append(deck[:randCardIndex], deck[randCardIndex+1:]...)
	
		randCardIndex = rand.Intn(len(deck))
		computerCards[deck[randCardIndex]], _ = strconv.Atoi(string(deck[randCardIndex][0]))
		deck = append(deck[:randCardIndex], deck[randCardIndex+1:]...)
	}

	fmt.Println("Your cards: ")
	for k, v := range playerCards {
		fmt.Println(k)
		playerPoints += v
	}
	fmt.Printf("Your points: %d\n\n", playerPoints)

	fmt.Println("The computer's cards")
	for k, v := range computerCards {
		fmt.Println(k)
		computerPoints += v
	}
	fmt.Printf("The computer's points: %d\n\n", computerPoints)

	if playerPoints == computerPoints {
		fmt.Println("Draw\nAnother round will start.\n")
		return Blackjack()
	}

	return checkForBlackJackWinner(playerPoints, computerPoints)
}

func shuffleDeck() []string {
	deck := []string{}
	for _, cardType := range []string{"clubs", "diamonds", "hearts", "spades"} {
		for i := 1; i <= 10; i++ {
			deck = append(deck, strconv.Itoa(i) + " of " + cardType)
		}
	}
	
	return deck
}

func checkForBlackJackWinner(playerPoints, computerPoints int) int {
 	switch {
	case playerPoints > 21 && computerPoints > 21:
		if playerPoints > computerPoints {
			fmt.Println("You lost.\n")
			return 0
		}

		fmt.Println("You won!\n")
		return 1
	case playerPoints < 21 && computerPoints < 21:
		if playerPoints > computerPoints {
			fmt.Println("You won!\n")
			return 1
		}

		fmt.Println("You lost.\n")
		return 0
	default:
		if playerPoints > computerPoints {
			fmt.Println("You lost.\n")
			return 0
		}

		fmt.Println("You won!\n")
		return 1
	}
}
