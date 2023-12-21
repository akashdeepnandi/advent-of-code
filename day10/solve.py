grid = open(0).read().strip().splitlines()
sr, sc = 0, 0
ROWS, COLS = len(grid), len(grid[0])

HAS_DOWNLINK = "S|F7"
HAS_UPLINK = "S|JL"
HAS_LEFTLINK = "S-7J"
HAS_RIGHTLINK = "S-FL"

for r in range(ROWS):
    found = False
    for c in range(COLS):
        if grid[r][c] == "S":
            sr, sc = r, c
            found = True
            break
    if found:
        break


def valid_pos(r, c):
    return 0 <= r < ROWS and 0 <= c < COLS


def get_neigbours(grid, pr, pc):
    directions = [
        (-1, 0),  # up
        (1, 0),  # down
        (0, 1),  # right
        (0, -1),  # left
    ]
    moves = [
        (HAS_UPLINK, HAS_DOWNLINK),  # for move up - source and neigbour
        (HAS_DOWNLINK, HAS_UPLINK),  # for move down - source and neigbour
        (HAS_RIGHTLINK, HAS_LEFTLINK),  # for move left
        (HAS_LEFTLINK, HAS_RIGHTLINK),  # for move right
    ]
    source = grid[pr][pc]

    neighbours = []
    for i, (dr, dc) in enumerate(directions):
        possible_source, possible_neigbour = moves[i]
        nr = pr + dr
        nc = pc + dc
        if valid_pos(nr, nc):
            if source in possible_source and grid[nr][nc] in possible_neigbour:
                neighbours.append((nr, nc))
    return neighbours


def dfs_find_path(grid, start):
    path = []
    stack = [start]
    visited = set()
    visited.add(start)
    while len(stack) > 0:
        current = stack.pop()
        path.append((current))
        r, c = current
        for neighbour in get_neigbours(grid, r, c):
            if neighbour not in visited:
                visited.add(neighbour)
                stack.append(neighbour)
    return path


path = dfs_find_path(grid, (sr, sc))
print(len(path) // 2)
