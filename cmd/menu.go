package main

import "fmt"

func runMenu() {
	parkingLot := make([]string, 10)
	var choice int

	for {
		fmt.Println("---Parking Lot Menu---")
		fmt.Println("1) Park a car")
		fmt.Println("2) Find a car")
		fmt.Println("3) Show parking lot (debug)")
		fmt.Println("4) Remove a car from parking lot")
		fmt.Println("5) Exit")

		fmt.Print("Please enter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input, please enter a number 1-5")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if choice == 5 {
			fmt.Println("👋 Goodbye!")
			break // this stops the loop and exits our program
		}

		handleAction(choice, &parkingLot)
	}

}