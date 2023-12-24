from functools import reduce
import sys

# python oneliner - using black formatter, so broke the line. For both part 1 and 2
ra = reduce(
    lambda r, d: (
        r[0]
        + 2 * (d[0] * d[1] + d[1] * d[2] + d[2] * d[0])
        + min(d[0] * d[1], d[1] * d[2], d[2] * d[0]),
        r[1] + d[0] * d[1] * d[2] + 2 * min(d[0] + d[1], d[1] + d[2], d[2] + d[0]),
    ),
    [tuple(map(int, line.split("x"))) for line in open(sys.argv[1])],
    (0, 0),
)
print(ra)

p, r = 0, 0
for line in open(sys.argv[1]):
    l, w, h = list(map(int, line.split("x")))
    p += 2 * (l * w + w * h + h * l) + min(l * w, w * h, h * l)
    r += l * w * h + 2 * min(l + w, w + h, h + l)
print("Total Paper", p)
print("Total Wrapping", r)
