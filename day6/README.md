# TOC

- [TOC](#toc)
- [About](#about)
    - [Advent of Code 2023 Day 2 --- Day 6: Wait For It ---](#advent-of-code-2023-day-2-----day-6-wait-for-it----)
        - [Solutions](#solutions)


# About

## Advent of Code 2023 Day 2 --- Day 6: Wait For It ---

### Solutions
- [Golang](main.go)
- [JavaScript](index.js)
- [Python](solve.py)

## Golang Performance

Evidently from the performance tests, the Quadratic solution is 99.8% faster.
```
goos: linux
goarch: amd64
pkg: github.com/akashdeepnandi/advent-of-code/day6
cpu: AMD Ryzen 7 6800H with Radeon Graphics

BenchmarkQuadratic-16              57334             21210 ns/op
BenchmarkQuadratic-16              56179             21408 ns/op
BenchmarkQuadratic-16              54760             21120 ns/op

BenchmarkBruteForce-16                78          13701323 ns/op
BenchmarkBruteForce-16                84          13992336 ns/op
BenchmarkBruteForce-16                80          14277573 ns/op
PASS
ok      github.com/akashdeepnandi/advent-of-code/day6   7.666s
```
