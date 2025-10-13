# Lookup function pseudocode - Checks if a car is already parked

This function checks whether a car (by its plate number) is already in the parking lot (hash table).
It uses the same hash function to find the starting index and then linearly probes forward if there's a collision.

```
DEFINE lookup(number_plate, parking_lot):
    # Step 1: Compute initial parking index
    index = hash(number_plate, parking_lot)
    start = index

    # Step 2: Search for the car using linear probing
    WHILE parking_lot[index] is not empty:
        IF parking_lot[index] is equal to parking_lot:
            RETURN true
        
        # Move to the next slot
        index = (index + one) MOD len of parking_lot

        # If we've looped back to the start, stop searching
        IF index == start:
            BREAK

    # Step 3: Car not found
    RETURN false