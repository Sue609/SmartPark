# Hash Function Pseudocode — Converts key → index

The hash function takes a plate (e.g., a license plate string) and table_size, and converts it into an integer index within the table size using the modulo operation. This ensures each key maps to a valid bucket.

```
DEFINE hash function (number_plate, parking_lot):
    sum = 0
    FOR each char IN number_plate:
        sum + ASCII of char
    RETURN sum MOD parking_lot