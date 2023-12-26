import sys

lines = open(sys.argv[1]).read().strip().splitlines()

# read line by line
def parse_instruction(s):
    m, si, ei = -1, 2, 4
    if len(s) == 4:
        m, si, ei = 0, 1, 3
    elif s[1] == "on":
        m = 1
    else:
        m = 2
    return [m] + list(map(int, s[si].split(","))) + list(map(int, s[ei].split(",")))


def operate_light(v, mode):
    if mode == 0:
        return 0 if v == 1 else 1
    elif mode == 1:
        return 1
    else:
        return 0


def operate_brightness(v, mode):
    if mode == 0:
        return v + 2
    elif mode == 1:
        return v + 1
    else:
        return max(0, v - 1)


def operate_decoration(lines, operate):
    grid = [[0 for _ in range(1000)] for _ in range(1000)]

    for line in lines:
        mode, start_r, start_c, end_r, end_c = parse_instruction(line.split())
        for r in range(start_r, end_r + 1):
            for c in range(start_c, end_c + 1):
                grid[r][c] = operate(grid[r][c], mode)

    count = 0
    for r in range(1000):
        for c in range(1000):
            count += grid[r][c]

    return count


print("Part 1:", operate_decoration(lines, operate_light))
print("Part 2:", operate_decoration(lines, operate_brightness))
# read all lines
