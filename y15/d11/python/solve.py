import sys

# read all lines
password = open(sys.argv[1]).read().strip().splitlines()[0]
itol = list("abcdefghijklmnopqrstuvwxyz")
ltoi = {l: i for i, l in enumerate(itol)}

invalid_letters = ["i", "l", "o"]


def increment_password(password: str):
    result = list(password)

    # find the left most invalid character
    for i, current_letter in enumerate(password):
        if current_letter in invalid_letters:
            previous_letter = result[i]
            letter_index = ltoi[previous_letter]
            next_previous = itol[(letter_index + 1) % len(itol)]
            result[i] = next_previous

            for j in range(i + 1, len(result)):
                result[j] = "a"

            return "".join(result)

    # valid
    for current_index, current_letter in enumerate(reversed(result)):
        letter_index = ltoi[current_letter]
        next_letter = itol[(letter_index + 1) % len(itol)]
        # handling the indexing as the current_index is reverse
        result[len(result) - (current_index + 1)] = next_letter
        if next_letter != "a":
            break

    return "".join(result)


def validate_password(password):
    three_straight = False
    last_letter = password[0]
    last_index = ltoi[last_letter]
    count = 1
    pair_count = {last_letter: 1}

    for i in range(1, len(password)):
        current_letter = password[i]
        current_index = ltoi[current_letter]

        # Checks invalid letters
        if current_letter in invalid_letters:
            return False

        # Checks three straight letters
        if count == 3:
            three_straight = True
        elif last_index + 1 == current_index:
            count += 1
        else:
            count = 1

        # Collects pairs or more
        if last_letter == current_letter:
            pair_count[current_letter] += 1
        else:
            pair_count[current_letter] = 1

        last_letter = current_letter
        last_index = current_index

    has_double_pairs = False
    pairs = 0
    for v in pair_count.values():
        if v >= 2:
            pairs += 1
        if pairs >= 2:
            has_double_pairs = True
            break

    return three_straight and has_double_pairs


for i in range(10000000):
    password = increment_password(password)
    if validate_password(password):
        break

print("Part 1:", password)

for i in range(10000000):
    password = increment_password(password)
    if validate_password(password):
        break

print("Part 2:", password)
