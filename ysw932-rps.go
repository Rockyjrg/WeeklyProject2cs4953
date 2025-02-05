package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Boss struct {
	name string
	move func(string) string
}

// function for game bosses
func stubbornBoss(userChoice string) string {
	return "rock"
}

func ghostBoss(userChoice string) string {
	choices := []string{"rock", "paper"}
	return choices[rand.IntN(len(choices))]
}

func demonBoss(userChoice string) string {
	//1 in 5 chance
	if rand.IntN(5) == 0 {
		return userChoice
	}
	choices := []string{"rock", "paper", "scissors"}
	return choices[rand.IntN(len(choices))]
}

// function to get random bosses
func getRandomBoss() Boss {
	bosses := []Boss{
		{"Stubborn Boss", stubbornBoss},
		{"Ghost Boss", ghostBoss},
		{"Demon Boss", demonBoss},
	}
	return bosses[rand.IntN(len(bosses))]
}

func fightBoss(boss Boss, playerHealth *int) bool {
	var bossHealth int = 3
	fmt.Printf("You are fighting %s!\n", boss.name)

	for *playerHealth > 0 && bossHealth > 0 {
		var userChoice string
		for {
			fmt.Println("It is your move. Choose rock, paper, or scissors.")
			fmt.Scan(&userChoice)
			userChoice = strings.ToLower(userChoice)

			if userChoice == "exit" {
				fmt.Println("You ran away, you live to fight again...")
				return true
			} else if userChoice == "rock" || userChoice == "paper" || userChoice == "scissors" {
				break //valid choice
			} else {
				fmt.Println("Invalid choice. Please enter rock, paper, or scissors.")
			}
		}

		var aiChoice string = boss.move(userChoice)
		fmt.Printf("%s chose: %s\n", boss.name, aiChoice)

		if userChoice == aiChoice {
			fmt.Println("It's a tie!")
		} else if (userChoice == "rock" && aiChoice == "scissors") ||
			(userChoice == "paper" && aiChoice == "rock") ||
			(userChoice == "scissors" && aiChoice == "paper") {
			fmt.Println("You landed a hit on the boss!")
			bossHealth--
		} else {
			fmt.Println("You got hit!")
			*playerHealth--
		}
		fmt.Printf("Your health: %d | %s health: %d\n", *playerHealth, boss.name, bossHealth)
	}

	if *playerHealth > 0 {
		fmt.Println("You have defeated the boss!")
	} else {
		fmt.Println("You have been defeated...")
	}

	return false //player didn't type exit
}

func main() {
	fmt.Println("Welcome to RPS dungeon crawl. Type exit to run away")

	var playerHealth int = 5

	for playerHealth > 0 {
		boss := getRandomBoss()
		//if the player fightboss returns true exit the game
		if fightBoss(boss, &playerHealth) {
			break
		}

		if playerHealth > 0 {
			fmt.Println("Moving to the next room...")
		}
	}
	fmt.Print("Game over. Thanks for playing.")
}
