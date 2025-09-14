import sys

graph: dict[str, dict[str, int]] = dict()
nodes: list[str] = []
for line in open(sys.argv[1]):
    parts = line.split()
    u, v, d = parts[0], parts[2], int(parts[-1])
    if u not in graph:
        graph[u] = dict()
    if v not in graph:
        graph[v] = dict()
    graph[u][v] = d
    graph[v][u] = d

    if u not in nodes:
        nodes.append(u)
    if v not in nodes:
        nodes.append(v)


def get_dist(u: str, v: str):
    return graph[u][v]


def permute(arr: list[str], i: int, result: list[list[str]]):
    if i == len(arr):
        result.append(arr[:])
        return

    for j in range(i, len(arr)):
        arr[i], arr[j] = arr[j], arr[i]
        permute(arr, i + 1, result)
        # restore
        arr[i], arr[j] = arr[j], arr[i]


def get_route_dist(route: list[str]) -> int:
    route_dist = 0
    for i in range(len(route) - 1):
        u = route[i]
        v = route[i + 1]
        route_dist += get_dist(u, v)
    return route_dist


routes: list[list[str]] = []

permute(nodes, 0, routes)
route_distances = [get_route_dist(r) for r in routes]

part1 = min(route_distances)

print("Part 1:", part1)

part2 = max(route_distances)

print("Part 2:", part2)
