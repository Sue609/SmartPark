# printLot function pseudocode

This function prints the current state of the parking lot, showing which slots are occupied or empty.
It prints the visualization of each slot and it's state.
It is used for debugging and monitoring the system's current parking lot status

```
DEFINE printLot(parking_lot):
    PRINT "Parking lot state"
    FOR each index i and value v in parking_lot:
        IF v is empty:
            PRINT "[i] empty"
        ELSE:
            PRINT "[i] v"