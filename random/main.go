package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// fmt.Println(rand.Intn(101))     // [0, 101)
	// fmt.Println(rand.Intn(100) + 1) // [1, 100]

	// val := rand.New(rand.NewSource(42))

	// fmt.Println(val.Intn(101)) // 97 for every run

	// fmt.Println(rand.Float32()) // between 0.0 and 1.0

	// Dice game
	for {
		// Show the menu
		fmt.Println("Welcome the game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2")
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		fmt.Printf("Die1 = %d and Die2 = %d, Total = %d\n\n", die1, die2, die1+die2)
	}

}
