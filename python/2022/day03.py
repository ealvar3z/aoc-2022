PRIORITY = ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


def part1(f):
    """
    Find the item type that appears in both compartments of each rucksack.
    What is the sum of the priorities of those item types?
    """
    sum = 0
    for line in f:
        mid = len(line) // 2
        fst, snd = line[:mid], line[mid:]
        # find what is common in both sets (i.e. the intersection)
        common = set(fst) & set(snd)
        sum += PRIORITY.index(next(iter(common)))

    return sum


def part2(file):
    """
    Find the item type that corresponds to the badges of each three-Elf group.
    What is the sum of the priorities of those item types?
    """
    sum = 0
    item = iter(file.read().split())

    while True:
        try:
            fst, snd, third = next(item), next(item), next(item)
        except StopIteration:
            return sum

        badge = set(fst) & set(snd) & set(third)
        sum += PRIORITY.index(next(iter(badge)))
