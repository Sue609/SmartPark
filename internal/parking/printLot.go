package parking
import "fmt"

func PrintLot(parkingLot [] string) {
	// Prints the current state of the parking lot

	fmt.Println("Parking lot state")

	for i, v := range parkingLot {
		if v == "" {
			fmt.Printf("[%d] empty\n", i)
		} else {
			fmt.Printf("[%d] %s\n", i, v)
		}
	}
}