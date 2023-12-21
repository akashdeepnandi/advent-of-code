
move up
things that can take move down
so if current is SLJ|
and up is this SF7|

move down
current is this SF7|
down is  SLJ|

move left

source L-FS
neigbour 7J-S

move right
source 7J-S
neigbour L-FS

This is a grid forms a loop
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...

With following rules:
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position

This grid forms a loop with these rules. Which i can find with dfs.
But the points that are inside the loop. What algorithms can i use?
