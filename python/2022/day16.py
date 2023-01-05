import math
import re
import sys
from collections import defaultdict
from functools import cache

RE = r"^Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([\w, ]+)$"

def part_two(f):
    valves = {}
    dist = defaultdict(lambda: defaultdict(lambda: math.inf))

    for i, flow_rate, tunnels in re.findall(RE, f.read(), re.MULTILINE):
        valves[i] = int(flow_rate)
        dist[i][i] = 0
        for j in tunnels.split(", "):
            dist[i][j] = 1

    for k in valves:
        for i in valves:
            for j in valves:
                dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

    @cache
    def backtrack(i, t, rem, elephant):
        answer = backtrack("AA", 26, rem, False) if elephant else 0
        for j in rem:
            if (next_t := t - dist[i][j] - 1) >= 0:
                answer = max(answer, valves[j] * next_t + backtrack(j, next_t, rem - {j}, elephant))
        return answer

    return backtrack("AA", 26, frozenset(x for x in valves if valves[x] > 0), True)

