# Insert function pseudocode - Park the car in the HASH TABLE
Takes in a car's number plate and table. First look at the IF the car is already parked using the lookup function.
If not we find the correct index using the hash function which uses **linear probing** to find an empty spot.
```
DEFINE (plate,table):
    # Step 1: Check if car already parked
    IF lookup(plate, table):
        PRINT "Car is already parked"
        RETURN
    
    # Step 2: Compute initial parking index
    index = hash(plate, table)
    start = index

    # Step 3: Linear probing - move forward until an empty slot is found
    WHILE plate[index] is occupied and table[index] != DELETED:
        MOVE to next slot index, index = (index + 1) MOD len(table)

        # Check if all spots are probed
        IF index == start:
            PRINT "Parking lot is full"
            RETURN

    # Step 4: Park the car
    table[index] = plate
    PRINT "Car parked successfully"

    
        