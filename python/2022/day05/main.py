from collections import deque
from re import findall


def part1(file):
    crates, directions = file.read().split("\n\n")
    stacks = []

    for c in crates.splitlines():
        for i, v in enumerate(range(1, len(c), 4)):
            print(c)
            while i >= len(stacks):
                stacks.append(deque())
            if c[v] != " ":
                stacks[i].append(c[v])

    for d in directions.splitlines():
        a, b, c = map(int, findall(r"\d+", d))
        for i in range(a):
            stacks[c - 1].appendleft(stacks[b - 1].popleft())

    return "".join(x[0] for x in stacks)


def part2(file):
    crates, directions = file.read().split("\n\n")
    stacks = []

    for c in crates.splitlines():
        for i, v in enumerate(range(1, len(c), 4)):
            print(c)
            while i >= len(stacks):
                stacks.append(deque())
            if c[v] != " ":
                stacks[i].append(c[v])

    for d in directions.splitlines():
        a, b, c = map(int, findall(r"\d+", d))
        temp = deque()
        for i in range(a):
            temp.appendleft(stacks[b - 1].popleft())
        stacks[c - 1].extendleft(temp)

    return "".join(x[0] for x in stacks)

