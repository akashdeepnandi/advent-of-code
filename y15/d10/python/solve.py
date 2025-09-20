import sys

lines = open(sys.argv[1]).read().strip().splitlines()
aoc_input = lines[0]


def look_and_say(num: str) -> str:
    result: list[str] = []
    current_digit = num[0]
    count = 0
    for d in num:
        if current_digit == d:
            count += 1
        else:
            result.append(str(count) + current_digit)
            current_digit = d
            count = 1
    # fill the last
    result.append(str(count) + current_digit)
    return "".join(result)


def evolve_sequence(num: str, generation: int) -> str:
    for _ in range(generation):
        num = look_and_say(num)
    return num


# Part 1
part1 = evolve_sequence(aoc_input, 40)
print("Part 1:", len(part1))

# Part 2
part2 = evolve_sequence(aoc_input, 50)
print("Part 2:", len(part2))
