import math
import operator
import re
from functools import cache

class T(tuple):
    def put(self, idx, val):
        return T(val if i == idx else x for i, x in enumerate(self))

    def zip_map(self, other, func):
            return T(func(x, y) for x, y in zip(self, other))

    def __add__(self, other):
        return self.zip_map(other, operator.add)

    def __sub__(self, other):
        return self.zip_map(other, operator.sub)
    
    def __mul__(self, other):
        return T(x * other for x in self)


def blueprints(f):
    ints = re.findall(r"\d+", f.read())
    _blueprints = []
    for offset in range(0, len(ints), 7):
        i, ore, clay, obs_ore, obs_clay, geode_ore, geode_obs = map(int, ints[offset: offset + 7])
        bp = T((ore, 0, 0, 0)), T((clay, 0, 0, 0)), T((obs_ore, obs_clay, 0, 0)), T((geode_ore, 0, geode_obs, 0))
        _blueprints.append((i, bp))
    return _blueprints

def req_time(cost, gain):
    if cost == 0:
        return 0
    if gain == 0:
        return math.inf
    return max(math.ceil(cost / gain), 0)

def solve(time, blueprint):
    curr = 0

    @cache
    def backtrack(time, blueprint, bots, inventory):
        nonlocal curr
        
        answer = inventory[3] + bots[3] * time
        curr = max(answer, curr)

        if answer + time * (time+1) // 2 <= curr:
            return answer

        for i in range(3):
            max_used = time * max(c[i] for c in blueprint) - bots[i] * (time -1)
            if inventory[i] > max_used:
                return backtrack(time, blueprint, bots, inventory.put(i, max_used))

        for idx, cost in enumerate(blueprint):
            earnings = cost - inventory
            time_needed = max(earnings.zip_map(bots, req_time)) + 1

            if idx < 3 and bots[idx] >= max(c[idx] for c in blueprint):
                continue

            if time - time_needed >= 0:
                new_inventory = inventory + bots * time_needed - cost
                new_bots = bots + [idx == i for i in range(4)]
                answer = max(answer, backtrack(time - time_needed, blueprint, new_bots, new_inventory))
                curr = max(answer, curr)

        return answer

    return backtrack(time, blueprint, T((1, 0, 0, 0)), T((0, 0, 0, 0)))


def part_one(f):
    return sum(i * solve(24, blueprint) for i, blueprint in blueprints(f))

def part_two(f):
    return math.prod(solve(32, blueprint) for i, blueprint in blueprints(f)[:3])
