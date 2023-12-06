INPUT = "./sinput" if False else "input"
sum1 = 0
sum2 = 0
### SOLUTION -- START
cardCopyTracker = []

def filterInt(nums):
    filtered = []
    for n in nums:
        if n.strip() != "":
            filtered.append(int(n))
    return filtered


with open(INPUT, "r") as f:
    lines = f.read().splitlines()
    # just learnt this technique
    cardCopyTracker = [1] * len(lines)
    for lineIdx, line in enumerate(lines):
        [winNums, myNums] = [filterInt(nums.split(" ")) for nums in line.split(": ")[1].split(" | ")]
        matchCount = 0
        for n in myNums:
            if n in winNums:
                matchCount += 1
        
        if matchCount > 0:
            # FOR PART 1
            points = pow(2, matchCount - 1)
            sum1 += points

            # FOR PART 2
            startIdx = lineIdx + 1
            endIdx = lineIdx + matchCount
            currentCardCopies = cardCopyTracker[lineIdx]
            for i in range(startIdx, endIdx + 1):
                cardCopyTracker[i] += currentCardCopies
    sum2 = sum(cardCopyTracker)

print("Part 1 Sum", sum1)
print("Part 2 Sum", sum2)

### SOLUTION -- END
