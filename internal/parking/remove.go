package parking
import "fmt"

func Remove(numberPlate string, parkingLot []string) {
	// Compute initial parking index of the car
	length := len(parkingLot)
	index := Hash(numberPlate, length)
	start := index

	// Search for the car using linear probing
	for parkingLot[index] != "" {
		if parkingLot[index] == numberPlate {
			parkingLot[index] = "DELETED"
			fmt.Println("The car was removed successfully: ", numberPlate)
			return
		}

		index = (index + 1) % length

		// If we have looped back to the start, stop searching
		if index == start {
			break
		}
	}

	// Car not found
	fmt.Println("Car not found: ", numberPlate)

}