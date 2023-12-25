import sys
from vector import Vector

moves = open(sys.argv[1]).read().strip()

move_dir = {
    "^": Vector(-1, 0),
    "v": Vector(1, 0),
    "<": Vector(0, -1),
    ">": Vector(0, 1),
}

santa_pos, both_pos = Vector(0, 0), [Vector(0, 0), Vector(0, 0)]
visited1, visited2 = set(), set()

for i, move in enumerate(moves):
    santa_pos = santa_pos + move_dir[move]
    visited1.add(str(santa_pos))  # for part 1

    both_pos[i % 2] = both_pos[i % 2] + move_dir[move]
    visited2.add(str(both_pos[i % 2]))  # for part 2

print("Part 1", len(visited1))
print("Part 2", len(visited2))
