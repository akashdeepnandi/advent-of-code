import sys

instructions = {}
for line in open(sys.argv[1]):
    opr, id = line.strip().split(" -> ")
    opr = opr.split()
    instructions[id] = (
        ""
        if len(opr) == 1
        else ("NOT" if len(opr) == 2 else opr[1]),  # operator if any
        opr[1] if len(opr) == 2 else opr[0],  # operand 1
        opr[2] if len(opr) == 3 else -1,  # operand 2 if any
    )

operators = {
    "AND": int.__and__,
    "OR": int.__or__,
    "LSHIFT": lambda x, y: x << y & 0xFFFF,
    "RSHIFT": int.__rshift__,
}


def find_wire_val(instructions, target):
    values = {}

    def get_wire_val(wire: str) -> int:
        if wire.isdigit():
            return int(wire)
        else:
            return values[wire] if wire in values else -1

    while True:
        resolved = True
        for tid in instructions:
            if tid in values:  # skip if the value has been already found
                continue
            if target in values:  # skip if the target value is already found
                resolved = True
                break
            gate, w1, w2 = instructions[tid]
            v1 = get_wire_val(w1)
            if v1 == -1:
                resolved = False
                continue
            if gate in ["", "NOT"]:  # single operand
                if gate == "NOT":
                    v1 = 65535 - v1
                values[tid] = v1
            else:  # two operand
                v2 = get_wire_val(w2)
                if v2 == -1:
                    resolved = False
                    continue
                values[tid] = operators[gate](v1, v2)

        if resolved:
            break
    return values[target]


part1 = find_wire_val(instructions, "a")
print("Part 1", part1)
instructions["b"] = (instructions["b"][0], str(part1), instructions["b"][2])
part2 = find_wire_val(instructions, "a")
print("Part 2", part2)
