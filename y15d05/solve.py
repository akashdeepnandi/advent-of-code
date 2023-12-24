import re
import sys

lines = open(sys.argv[1]).read().strip().splitlines()
# Part 1
count = 0
for line in lines:
    vowels = re.findall(r"(a|e|i|o|u)", line)
    if len(vowels) < 3:
        continue
    doubles = re.findall(r"(\w)\1", line)
    if len(doubles) <= 0:
        continue
    bad = re.findall(r"(ab|cd|pq|xy)", line)
    if len(bad) > 0:
        continue
    count += 1

print(count)

# Part 2
count = 0
for line in lines:
    pairs = re.findall(r"(\w{2}).*\1", line)
    if len(pairs) <= 0:
        continue
    one_between = re.findall(r"(\w)\w\1", line)
    if len(one_between) <= 0:
        continue
    count += 1

print(count)
