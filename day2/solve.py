import re

INPUT = "./input"
sum1 = 0
sum2 = 0
### SOLUTION -- START

cubeLimit = {"red": 12, "green": 13, "blue": 14}

with open(INPUT, "r") as f:
    lines = f.read().splitlines()
    for lineIdx, line in enumerate(lines):
        # get the cube counts
        redCounts = [int(m) for m in re.findall(r"(\d+) red", line)]
        greenCounts = [int(m) for m in re.findall(r"(\d+) green", line)]
        blueCounts = [int(m) for m in re.findall(r"(\d+) blue", line)]

        # get the max count
        redMax = max(redCounts)
        greenMax = max(greenCounts)
        blueMax = max(blueCounts)

        ## FOR PART 1: check if max count is valid
        isGameValid = (
            redMax <= cubeLimit["red"]
            and greenMax <= cubeLimit["green"]
            and blueMax <= cubeLimit["blue"]
        )
        if isGameValid:
            sum1 += lineIdx + 1
        ## FOR PART 2: multiply all the max counts
        power = redMax * greenMax * blueMax
        sum2 += power

print("Part 1 Sum", sum1)
print("Part 2 Sum", sum2)

### SOLUTION -- END
