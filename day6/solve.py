import re
import math

INPUT = "./sinput" if True else "input"
sum1 = 0
sum2 = 0


### SOLUTION -- START
def getWinCountUsingRoots(time, distance):
    # here from ax2 + bx + c, a = 1, b = -time, c = distance
    sqrt = math.sqrt(time * time - 4 * distance)
    r1 = (time - sqrt) / 2
    r2 = (time + sqrt) / 2
    if r1.is_integer():
        r1 += 1
    if r2.is_integer():
        r2 -= 1
    r1 = math.ceil(r1)
    r2 = math.floor(r2)
    return (r2 - r1) + 1

def createBigInt(digits):
    d = ""
    for digit in digits:
        d += digit
    return int(d)

with open(INPUT, "r") as f:
    lines = f.read().splitlines()

    times, distances = [[int(m) for m in re.findall(r"(\d+)", l)] for l in lines]
    bigTime, bigDistance = [createBigInt(re.findall(r"(\d+)", l)) for l in lines]

    part1Sol = 1
    for i, t in enumerate(times):
        d = distances[i]
        wins = getWinCountUsingRoots(t, d)
        part1Sol *= wins
    part2sol = getWinCountUsingRoots(bigTime, bigDistance)

    print("Part 1", part1Sol)
    print("Part 2", part2sol)

### SOLUTION -- END
