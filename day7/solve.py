from typing import Counter


INPUT = "./sinput" if True else "input"
sum1 = 0
sum2 = 0
### SOLUTION -- START

cardMap1 = {"T": "A", "J": "B", "Q": "C", "K": "D", "A": "E"}


def getHandRank1(h):
    counts = sorted(list(Counter(h).values()))
    h = "".join([cardMap1.get(card, card) for card in h])
    x = strength(counts, h)
    return strength(counts, h)


cardMap2 = {"T": "A", "J": "1", "Q": "C", "K": "D", "A": "E"}


def getHandRank2(h):
    h = "".join([cardMap2.get(card, card) for card in h])
    cardCounter = Counter(h)
    sortedCount = sorted(cardCounter.items(), key=lambda c: c[1], reverse=True)
    if len(sortedCount) > 1:
        fc, sc = sortedCount[0][0], sortedCount[1][0]
        highestCard = sc if fc == "1" else fc

        if "1" in cardCounter and highestCard != ".":
            cardCounter[highestCard] += cardCounter["1"]
            del cardCounter["1"]

    counts = sorted(list(cardCounter.values()))

    return strength(counts, h)


def strength(counts, h):
    if counts == [5]:
        return (6, h)
    elif counts == [1, 4]:
        return (5, h)
    elif counts == [2, 3]:
        return (4, h)
    elif counts == [1, 1, 3]:
        return (3, h)
    elif counts == [1, 2, 2]:
        return (2, h)
    elif counts == [1, 1, 1, 2]:
        return (1, h)
    elif counts == [1, 1, 1, 1, 1]:
        return (0, h)


with open(INPUT, "r") as f:
    lines = f.read().splitlines()
    hands1 = []
    hands2 = []
    for lineIdx, line in enumerate(lines):
        h, b = line.split()[:2]
        hands1.append((getHandRank1(h), int(b)))
        hands2.append((getHandRank2(h), int(b)))
    hands1 = sorted(hands1, key=lambda x: x[0])
    hands2 = sorted(hands2, key=lambda x: x[0])

    sum1 = sum([(i + 1) * b for i, (_, b) in enumerate(hands1)])
    sum2 = sum([(i + 1) * b for i, (_, b) in enumerate(hands2)])

print("Part 1 Sum", sum1)
print("Part 2 Sum", sum2)

### SOLUTION -- END
