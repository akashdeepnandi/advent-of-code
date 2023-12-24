import sys
# read the line fro argument, parse it, make a character list
input = list(open(sys.argv[1]).read().strip().splitlines()[0])
f, pos = 0, 0
for i, p in enumerate(input):
    # for part 1, check if the char is ( then add 1 else -1
    f += 1 if p == "(" else -1
    # for part 2, check if we have reached floor -1 then capture the pos
    if f == -1 and pos == 0:
        pos  = i + 1

print("Floor", f)
print("Position", pos)
