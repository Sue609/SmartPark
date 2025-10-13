DEFINE hash function (plate, table_size):
    sum = 0
    for each char in plate:
        sum + ASCII of char
    return sum % table_size