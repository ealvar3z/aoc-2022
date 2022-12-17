from functools import total_ordering
from math import prod

def cmp(x, y):
    if isinstance(x, int) and isinstance(y, int):
        return x - y

    if isinstance(x, list) and isinstance(y, list):
        for i, j in zip(x, y):
            if ans := cmp(i, j):
                return ans
        return len(x) - len(y)

    if isinstance(x, list):
        return cmp(x, [y])

    if isinstance(y, list):
        return cmp([x], y)

    assert False

def part_one(f):
    """
    Determine which pairs of packets are already in the right order. What is
    the sum of the indices of those pairs?
    """
    pairs = [[eval(a) for a in pair.splitlines()] 
            for pair in f.read().split("\n\n")]
    return sum(i+1 for i, (x, y) in enumerate(pairs) if cmp(x, y) < 0)

@total_ordering
class Orderder:
    def __init__(self, x):
        self.x = x

    def __lt__(self, any):
        return cmp(self.x, any.x) < 0
    
    def __eq__(self, any):
        return cmp(self.x, any.x) == 0

def part_two(f):
    """
    Organize all of the packets into the correct order. What is the decoder key
    for the distress signal?
    To find the decoder key for this distress signal, you need to determine the
    indices of the two divider packets and multiply them together.
    """
    packets = [eval(p) for p in f.read().splitlines() if len(p) > 0]
    dividers = [[2]], [[6]]
    ans = sorted([*packets, *dividers], key=Orderder)
    return prod(ans.index(p) + 1 for p in dividers)
