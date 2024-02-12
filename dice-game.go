package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll() int {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1 // Generate random number between 1 and 6
}

func main() {
	var players int

	// Get valid player count
	for {
		fmt.Print("Enter the number of players (2 - 4): ")
		fmt.Scanln(&players)

		if players >= 2 && players <= 4 {
			break
		}
		fmt.Println("Must be between 2 - 4 players.")
	}

	maxScore := 50
	playerScores := make([]int, players)

	for max(playerScores) < maxScore {
		for playerIdx := range playerScores {
			fmt.Printf("\nPlayer number %d's turn!\n", playerIdx+1)
			fmt.Printf("Your total score is: %d\n", playerScores[playerIdx])
			var currentScore int

			for {
				var shouldRoll string
				fmt.Print("Would you like to roll (y/n)? ")
				fmt.Scanln(&shouldRoll)

				if shouldRoll[0] != 'y' && shouldRoll[0] != 'Y' {
					break
				}

				value := roll()
				if value == 1 {
					fmt.Println("You rolled a 1! Turn done.")
					currentScore = 0
					break
				} else {
					currentScore += value
					fmt.Printf("You rolled a: %d\n", value)
					fmt.Printf("Your score is: %d\n", currentScore)
				}
			}

			playerScores[playerIdx] += currentScore
			fmt.Printf("Your total score is: %d\n", playerScores[playerIdx])
		}
	}

	// Find and announce winner
	maxScore = max(playerScores)
	winningIdx := indexOf(playerScores, maxScore)
	fmt.Printf("Player number %d is the winner with a score of: %d\n", winningIdx+1, maxScore)
}

// Helper functions for finding max and index
func max(values []int) int {
	maxVal := values[0]
	for _, value := range values {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

func indexOf(values []int, target int) int {
	for i, value := range values {
		if value == target {
			return i
		}
	}
	return -1 // Not found
}
