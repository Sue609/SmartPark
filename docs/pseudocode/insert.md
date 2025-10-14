# Insert function pseudocode - Park the car in the HASH parking_lot
Takes in a car's number number_plate and parking_lot. First look at the IF the car is already parked using the lookup function.
If not we find the correct index using the hash function which uses **linear probing** to find an empty spot.
```
DEFINE Insert(number_plate, parking_lot):
    # Step 1: Check if car already parked
    IF lookup(number_plate, parking_lot):
        PRINT "Car is already parked"
        RETURN
    
    # Step 2: Compute initial parking index
    index = hash(number_plate, parking_lot)
    start = index

    # Step 3: Linear probing - move forward until an empty slot is found
    WHILE parking_lot[index] is NOT empty AND parking_lot[index] is NOT DELETED:
        MOVE to next slot index, index = (index + 1) MOD len(parking_lot)

        # Check if all spots are probed
        IF index == start:
            PRINT "Parking lot is full"
            RETURN

    # Step 4: Park the car
    parking_lot[index] = number_plate
    PRINT "Car parked successfully"

    
        