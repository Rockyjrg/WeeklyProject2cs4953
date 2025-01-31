package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	fmt.Println("We will be playing rock, paper, scissors!")

	for {
		rps := [3]string{"rock", "paper", "scissors"}
		randomChoice := rand.IntN(len(rps))
		randomElement := rps[randomChoice]
		var userChoice string
		fmt.Print("Ur move: ")
		fmt.Scan(&userChoice)
		fmt.Println("Ai chose: " + randomElement)
		if userChoice == "rock" {
			if randomElement == "scissors" {
				fmt.Println("You win!")
				continue
			} else if randomElement == "paper" {
				fmt.Println("You lose!")
				continue
			} else {
				fmt.Println("You tie!")
				continue
			}
		} else if userChoice == "paper" {
			if randomElement == "rock" {
				fmt.Println("You win!")
				continue
			} else if randomElement == "scissors" {
				fmt.Println("You lose!")
				continue
			} else {
				fmt.Println("You tie!")
				continue
			}
		} else if userChoice == "scissors" {
			if randomElement == "paper" {
				fmt.Println("You win!")
				continue
			} else if randomElement == "rock" {
				fmt.Println("You lose!")
				continue
			} else {
				fmt.Println("You tie!")
				continue
			}
		} else if userChoice == "exit" {
			fmt.Println("Ending game, bye bye!")
			break
		} else {
			fmt.Println("Please enter a valid answer.")
			continue
		}
	}
}
