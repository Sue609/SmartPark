package main

import (
	"fmt"
	"SmartPark/internal/parking"
)

// Connects the UI to our logic
func handleAction(choice int, parkingLot *[]string) {
	switch choice {
	case 1:
		fmt.Println("Enter the number plate: ")
		var numberPlate string
		fmt.Scanln(&numberPlate)
		parking.Insert(numberPlate, *parkingLot)

	case 2:
		fmt.Println("Enter the number plate to find")
		var numberPlate string
		fmt.Scanln(&numberPlate)
		if parking.Lookup(numberPlate, *parkingLot) {
			fmt.Println("✅ Car found!")
		} else {
			fmt.Println("❌ Car not found.")
		}

	case 3:
		parking.PrintLot(*parkingLot)

    case 4:
        fmt.Print("Enter the plate to remove: ")
        var numberPlate string
        fmt.Scanln(&numberPlate)
        parking.Remove(numberPlate, *parkingLot)

    default:
        fmt.Println("Please choose a number between 1 and 5.")
	}
}
