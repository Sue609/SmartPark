package parking

func hash(number_plate string, parking_lot []string) int {
	sum := 0

	for _, char := range number_plate {
		sum += int(char)
	}

	return sum % len(parking_lot)
}