grid = open(0).read().strip().splitlines()
ROWS = len(grid)
COLS = len(grid[0])

empty_rows = []
empty_cols = []
galaxies = []

# get empty rows and cols
for c in range(COLS):
    empty = True
    for r in range(ROWS):
        row = grid[r]
        if "#" not in row: # no # in row, empty
            empty_rows.append(r)
        if grid[r][c] == "#": # check the character is galaxy, put it in list
            galaxies.append((r, c))
            empty = False
    if empty: # if it passes the column empty check then put in empty cols
        empty_cols.append(c)

def get_min_max(a, b):
    return (b, a) if a > b else (a, b)

# common function to find the distance
def find_distance_sum(expansion_rate):
    sum = 0
    for i in range(len(galaxies) - 1):
        (r1, c1) = galaxies[i]
        for j in range(i + 1, len(galaxies)):
            (r2, c2) = galaxies[j]
            d = abs(r1 - r2) + abs(c1 - c2) # using manhattan distance
            minr, maxr = get_min_max(r1, r2)
            minc, maxc = get_min_max(c1, c2)
            for er in empty_rows:
                if er > minr and er < maxr:
                    d += expansion_rate - 1

            for ec in empty_cols:
                if ec > minc and ec < maxc:
                    d += expansion_rate - 1
            sum += d
    return sum

print("PART 1", find_distance_sum(2))
print("PART 2", find_distance_sum(1000000))
