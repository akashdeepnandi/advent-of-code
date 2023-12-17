import sys
from functools import cache

@cache
def get_arrangements(springs, groups):
    if springs == "":
        return 1 if groups == () else 0
    if groups == ():
        return 0 if "#" in springs else 1

    result = 0
    group = groups[0]
    current = springs[0]
    springs_l = len(springs)

    if current == "." or current == "?":  # start can be a dot
        result += get_arrangements(springs[1:], groups)
    if current == "#" or current == "?":  # start can be #, starting of a block
        if (
            group <= springs_l
            and "." not in springs[:group]
            and (group == springs_l or springs[group] != "#")
        ):
            result += get_arrangements(springs[group + 1 :], groups[1:])
    return result

total = 0
for line in open(sys.argv[1]):
    springs, groups = line.split()
    # FOR PART 2
    springs = "?".join([springs] * 5)
    groups = ",".join([groups] * 5)
    groups = tuple([int(g) for g in groups.split(",")])
    total += get_arrangements(springs, groups)

print(total)
