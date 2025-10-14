package parking

func Lookup(numberPlate string, parkingLot []string) bool {
	// Compute initial parking index
	length := len(parkingLot)
	index := Hash(numberPlate, length)
	start := index

	// Search for the car using linear probing
	for parkingLot[index] != "" {
		if parkingLot[index] == numberPlate {
			return true
		}

		index = (index + 1) % length

		if index == start {
			break
		}
	}
	return false

}