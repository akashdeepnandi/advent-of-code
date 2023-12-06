# TOC

- [TOC](#toc)
- [About](#about)
    - [Advent of Code 2023 --- Day 5: If You Give A Seed A Fertilizer ---](#advent-of-code-2023-----day-5-if-you-give-a-seed-a-fertilizer----)
        - [Solutions](#solutions)
            - [Working the problem](#working-the-problem)
    - [FIRST EPIC FAILED](#first-epic-failed)


# About

## Advent of Code 2023 --- Day 5: If You Give A Seed A Fertilizer ---

### Solutions
- [Golang](main.go)
- [JavaScript](index.js)
- [Python](solve.py)

#### Working the problem

First input line: 79 14 55 13

```
seed-to-soil map:
50 98 2
52 50 48
```

seed-to-soil:
`destRangeStart` `sourceRangeStart` `range length`

destRange = [50, 51]
sourceRangeStart = [98, 99]

With this we know,
98 maps to 50
99 maps to 51

so with this, we nbasical create the seed to soil map
similarly we need to create the other maps
then use that dat to find the lowest location number
50 98 2

## FIRST EPIC FAILED
MAXED OUT MY RAM
seed >= dstart && seed < dstart + range
mem=0 RSS=50752 elapsed=0:00.15 cpu.sys=0.09 user=0.15

