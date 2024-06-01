import json
import sys

# read line by line
p1 = 0
p2 = 0
for line in open(sys.argv[1]):
    original = line.rstrip()

    new = eval(original)
    p1 += len(original) - len(new)

    new = json.dumps(original)
    p2 += len(new) - len(original)

print("Part 1:", p1)
print("Part 2:", p2)
