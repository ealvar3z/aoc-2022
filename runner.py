from argparse import ArgumentParser
from time import time_ns
from traceback import print_exc
from datetime import datetime, timedelta, timezone
from importlib import import_module, reload


def main(func, filename="filename"):
    """runner function to speed up seeing the solution
    """
    try:
        with open(filename) as f:
            try:
                start = time_ns()
                print(func(f), end="\t")
                end = time_ns()
                print(f"[{(end - start) / 10**6:.3f} ms]")
            except:
                print_exc()
    except FileNotFoundError:
        print()


if __name__ == "__main__":
    now = datetime.now(timezone(timedelta(hours=-5)))
    p = ArgumentParser(description="AoC Runner")
    p.add_argument("--year", "-y", type=int,
                   help="which year", default=now.year)
    p.add_argument("--day", "-d", type=int, help="which day", default=now.day)
    args = p.parse_args()

    try:
        from aocd import get_data
    except ImportError:
        pass
    else:
        with open(f"input/{args.year}/day{args.day:02}.txt", "w") as f:
            f.write(get_data(day=args.day, year=args.year))

    module_name = f"python.{args.year}.day{args.day:02}"
    print(f"{module_name}")

    module = import_module(module_name)

    for i in ("part1", "part2"):
        if not hasattr(module, i):
            continue
        print(f"--- {i} ---")
        print("sample:", end="\t")
        main(getattr(module, i),
             f"input/{args.year}/day{args.day:02}_sample.txt")
        reload(module)
        print("actual:", end="\t")
        main(getattr(module, i), f"input/{args.year}/day{args.day:02}.txt")
