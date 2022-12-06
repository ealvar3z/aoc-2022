"""
Every section has a unique ID number,  and each Elf is assigned a range of
section IDs. To try to quickly find overlaps and reduce duplicated effort, the
Elves pair up and make a big list of the section assignments for each pair
(your puzzle input).

For example, consider the following list of section assignment pairs:

2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8

For the first few pairs, this list means:

    Within the first pair of Elves, the first Elf was assigned sections 2-4
    (sections 2, 3, and 4), while the second Elf was assigned sections 6-8
    (sections 6, 7, 8).  The Elves in the second pair were each assigned two
    sections.  The Elves in the third pair were each assigned three sections:
    one got sections 5, 6, and 7, while the other also got 7, plus 8 and 9.

In how many assignment pairs does one range fully contain the other?
"""


def part1(file):
    from aocd import data
    from parse import parse
    count = 0
    for line in data.splitlines():
        fst, snd, third, fourth = parse(
            "{:d}-{:d},{:d}-{:d}", line)
        if fst > third or (fst == third and snd < fourth):
            fst, snd, third, fourth = third, fourth, fst, snd
        count += fourth <= snd
    return count


def part2(file):
    from aocd import data
    from parse import parse
    count = 0
    for line in data.splitlines():
        fst, snd, third, fourth = parse(
            "{:d}-{:d},{:d}-{:d}", line)
        if fst > third or (fst == third and snd < fourth):
            fst, snd, third, fourth = third, fourth, fst, snd
        count += snd >= third
    return count
