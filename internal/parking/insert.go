package parking
import "fmt"

func Insert(numberPlate string, parkingLot [] string) {
	// Check if the car is already parked
	if Lookup(numberPlate, parkingLot) {
		fmt.Println("The car is already parked: ", numberPlate)
		return
	}

	length := len(parkingLot)
	index := Hash(numberPlate, length)
	start := index

	// move forward until an empty slot is found
	for parkingLot[index] != "" && parkingLot[index] != "DELETED" {
		index = (index + 1) % length

		if index == start {
			fmt.Println("Parking lot is full")
			return
		}
	}

	// park the car
	parkingLot[index] = numberPlate
	fmt.Println(numberPlate, "Parked successfully")
}