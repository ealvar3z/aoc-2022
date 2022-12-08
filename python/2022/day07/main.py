from collections import defaultdict
from pathlib import Path


def part_one(file):
    cwd = Path("/")
    dirs = defaultdict(int)

    for line in f.read().splitlines():
        match line.split():
            case ["$", "cd", newone]:
                cwd = cwd/newone
                cwd = cwd.resolve()
            case [sz, _] if sz.isdigit():
                sz = int(sz)
                for path in [cwd, list(cwd.parents)]:
                    dirs[path] += sz

    return sum(total for total in dirs.values() if total < 100000)


DISK_AVAIL = 70000000
UNUSED = 30000000
def part_two(f):
    cwd = Path("/")
    dirs = defaultdict(int)

    for line in f.read().splitlines():
        match line.split():
            case ["$", "cd", newone]:
                cwd = cwd/newone
                cwd = cwd.resolve()
            case [sz, _] if sz.isdigit():
                sz = int(sz)
                for path in [cwd, list(cwd.parents)]:
                    dirs[path] += sz

    return min(total for total in dirs.values()
               if dirs[Path("/")] - total <= DISK_AVAIL - UNUSED)
