import re

INPUT = "./input"
sum1 = 0
sum2 = 0
### SOLUTION -- START

def checkSymbol(s):
    return re.search(r"[^\d\.]", s)


def createLoc(x, y):
    return f"x:{x},y:{y}"


def checkIfGear(s, x, y):
    return (
        f"{createLoc(x, y)}"
        if s == "*"
        else ""
    )

with open(INPUT, "r") as f:
    lines = f.read().splitlines()
    allNumbers = []
    allGears = {}
    for lineIdx, line in enumerate(lines):
        numbers = [[int(m.group(0)), m.span()] for m in re.finditer(r"(\d+)", line)]
        for num in numbers:
            isPart = False
            [n, (mStart, mEnd)] = num
            startIdx = max(0, mStart - 1)
            endIdx = min(len(lines) - 1, mEnd)
            gearCoor = ""

            # CHECK ADJACENT LEFT
            symbol = checkSymbol(line[startIdx])
            if symbol:
                isPart = True
                gearCoor = checkIfGear(symbol.group(0), startIdx, lineIdx)

            # CHECK ADJACENT RIGHT
            symbol = checkSymbol(line[endIdx])
            if symbol:
                isPart = True
                gearCoor = checkIfGear(symbol.group(0), endIdx, lineIdx)

            # CHECK ADJACENT TOP ROW
            if lineIdx != 0:
                for sIdx in range(startIdx, endIdx + 1):
                    symbol = checkSymbol(lines[lineIdx - 1][sIdx])
                    if symbol:
                        isPart = True
                        gearCoor = checkIfGear(symbol.group(0), sIdx, lineIdx - 1)
                        break

            # CHECK ADJACENT BOTTOM ROW
            if lineIdx < len(lines) - 1:
                for sIdx in range(startIdx, endIdx + 1):
                    symbol = checkSymbol(lines[lineIdx + 1][sIdx])
                    if symbol:
                        if n == 114:
                            print(sIdx, startIdx, symbol.group(0))
                        isPart = True
                        gearCoor = checkIfGear(symbol.group(0), sIdx, lineIdx + 1)
                        break

            if isPart:
                sum1 += n
                if gearCoor:
                    numList = allGears.get(gearCoor)
                    numList = numList + [n] if isinstance(numList, list) else [n]
                    allGears[gearCoor] = numList
                    if len(numList) == 2:
                        [p0, p1] = numList
                        gearRatio = p0 * p1
                        sum2 += gearRatio


print("Part 1 Sum", sum1)
print("Part 2 Sum", sum2)

### SOLUTION -- END
