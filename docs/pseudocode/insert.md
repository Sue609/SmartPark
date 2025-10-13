# Insert function pseudocode
Takes in the plate and table, we first look at the IF the car is already parked in the parking lot using the "lookup function".
If the car is not parked, we find the index using the hash function, we traverse our hash map until we find an empty spot where we can safely park.

```
DEFINE (plate,table):
    # Step 1: Check is car already exists
    IF lookup(plate, table):
        PRINT "Car is already parked"
        RETURN
    
    # Step 2: Compute initial index
    index = hash(plate, table)
    start = index

    # Step 3: Linear probing - collision resolution technique
    WHILE plate[index] is occupied and table[index] != DELETED:
        MOVE to next slot index, index = (index + 1) MOD len(table)

        # Check if all spots are probed
        IF index == start:
            PRINT "Parking lot is full"
            RETURN

    # Step 4: Park the car
    table[index] = plate
    PRINT "Car parked successfully"

    
        