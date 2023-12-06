import re

INPUT = "./sinput" if True else "input"
sum1 = 0
sum2 = 0
### SOLUTION -- START

with open(INPUT, "r") as f:
    lines = f.read().splitlines()
    for lineIdx, line in enumerate(lines):
        print(line)

print("Part 1 Sum", sum1)
print("Part 2 Sum", sum2)

### SOLUTION -- END
