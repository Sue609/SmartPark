package parking

func Hash(number_plate string, table_size int) int {
	sum := 0

	for _, char := range number_plate {
		sum += int(char)
	}

	return sum % table_size
}