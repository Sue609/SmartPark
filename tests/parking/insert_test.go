package tests

import (
	"SmartPark/internal/parking"
	"testing"
)

func TestInsertCar(t *testing.T) {
	parkingLot := make([]string, 5)
	plate := "KCA123A"

	parking.Insert(plate, parkingLot)

	found := parking.Lookup(plate, parkingLot)

	if !found {
		t.Errorf("Expected car %s to be found after insertion", plate)
	}
}