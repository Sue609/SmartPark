# Main function pseudocode
It serves as the user interface that allows interaction with the parking lot through a menu-driven program.

From this main shell, users can:
- Park a car
- Find a car
- View the current parking lot status
- Remove parked car
- Exit the system

It continuously loops until the user chooses exit, handling invalid inputs gracefully and calling the core functions (insert, lookup, remove, and printLot) that manages the parking logic.


DEFINE main:
    # Step 1: Initialize parking lot with 10 slots
    CREATE an array of empty buckets 'parking_lot', size N
    DECLARE variable choice as integer

    LOOP FOREVER:
        PRINT "---Parking Lot Menu---"
        PRINT "1) Park a car"
        PRINT "2) Find a car"
        PRINT "3) Show parking lot (debug)"
        PRINT "4) Remove a car from parking lot"
        PRINT "5) Exit"
        PRINT "Enter choice: "

        # Step 2: Get user choice
        READ choice

        # Step 3: Validate input
        IF choice is not between 1 and 5:
            PRINT "Please enter a number 1–5"
            CONTINUE LOOP

        # Step 4: Perform operation based on choice
        SWITCH choice:
            CASE 1:
                PRINT "Enter the number plate of your car:"
                READ number_plate
                CALL insert(number_plate, parking_lot)

            CASE 2:
                PRINT "Enter the number plate to lookup:"
                READ number_plate
                IF lookup(number_plate, parking_lot) == true:
                    PRINT "Car found"
                ELSE:
                    PRINT "Car not found"

            CASE 3:
                CALL printLot(parking_lot)

            CASE 4:
                PRINT "Enter the number_plate to delete:"
                READ number_plate
                CALL remove(number_plate, parking_lot)

            CASE 5:
                PRINT "Bye"
                EXIT LOOP

            DEFAULT:
                PRINT "Invalid choice, pick between 1–5"



