import sys
# read line by line
for line in open(sys.argv[1]):
    print(line.strip())
# read all lines
lines = open(sys.argv[1]).read().strip().splitlines()
