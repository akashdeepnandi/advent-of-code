from math import inf
from heapq import heappush, heappop


grid = [list(map(int, line.strip())) for line in open(0)]
R = len(grid)
C = len(grid[0])
# h, r, c, dr, dc, n

def pqPop(q):
    minH = inf
    minI = -1
    for i, p in enumerate(q):
        h = p[0]
        if h < minH:
            minH = h
            minI = i

    return q.pop(minI)

q = [(0,0,0,0,0,0)]
seen = set()

i = 0
while len(q) > 0:
    i += 1
    # current = pqPop(q)
    current = heappop(q)
    h, r, c, dr, dc, n = current
    print(r, c)

    if r == R - 1 and c == C - 1:
        print(h)
        break

    if (r, c, dr, dc, n) in seen:
        continue
    seen.add((r, c, dr, dc, n))

    if n < 3 and (dr, dc) != (0, 0):
        nr = r + dr
        nc = c + dc
        if nr >= 0 and nr < R and nc >=0 and nc < C:
            # q.append((h + grid[nr][nc], nr, nc, dr, dc, n + 1))
            heappush(q, (h + grid[nr][nc], nr, nc, dr, dc, n + 1))

    for ndr, ndc in [(0, 1), (1, 0), (0, -1), (-1, 0)]:
        if (ndr, ndc) != (dr, dc) and (ndr, ndc) != (-dr, -dc):
            nr = r + ndr
            nc = c + ndc
            if nr >= 0 and nr < R and nc >=0 and nc < C:
                # q.append((h + grid[nr][nc], nr, nc, ndr, ndc, 1))
                heappush(q, (h + grid[nr][nc], nr, nc, ndr, ndc, 1))

    if i == 5:
        break
    # if i % 10000 == 0:
    #     print("calcuting...")
