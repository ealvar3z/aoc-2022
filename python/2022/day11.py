from collections import Counter, deque
from fileinput import input
from math import lcm
from operator import add, mul, mod
from re import compile, match

OP_RE = compile(r"^new = (?P<a>\w+) (?P<op>[\+\*]) (?P<b>\w+)$")
TEST_RE = compile(r"^(?P<op>\w+) by (?P<num>\w+)")
ACTION_RE = compile(r"^throw to monkey (?P<monkey>\w+)")
OP_DICT = {
    "+": add,
    "*": mul,
}
TEST_DICT = {
    "divisible": lambda x, y: mod(x, y) == 0,
}


def parse_ops(string):
    if pattern := match(OP_RE, string):
        def f(old):
            new = OP_DICT[pattern["op"]](
                old if pattern["a"] == "old" else int(pattern["a"]),
                old if pattern["b"] == "old" else int(pattern["b"]),
            )
            return new

        return f


def parse_one_tests(string, is_true, is_false):
    if pattern := match(TEST_RE, string):

        def f(value):
            test = TEST_DICT[pattern["op"]](value, int(pattern["num"]))
            return is_true if test else is_false

        return f


def parse_two_tests(string):
    if pattern := match(TEST_RE, string):
        return int(pattern["num"])

def parse_action(string):
    if pattern := match(ACTION_RE, string):
        return int(pattern["monkey"])


def whats_da_business(counter):
    return mul(*map(lambda p: p[1], counter.most_common(2)))


def parse_monkeys():
    lines = []
    for line in input(encoding="utf-8"):
        match line.strip():
            case "":
                yield lines
                lines = []
            case _ as line:
                lines.append(line)
    yield lines


def part_one():
    counter = Counter()
    monkeys = []
    for lines in parse_monkeys():
        monkey_dict = {
            "name": lines[0].strip(":").casefold(),
            "items": deque(
                map(int, lines[1].split(": ", maxsplit=1)[1].split(", "))
            ),
            "operation": parse_ops(lines[2].split(": ", maxsplit=1)[1]),
            "test": parse_tests(
                lines[3].split(": ", maxsplit=1)[1],
                lines[4].split(": ", maxsplit=1)[1],
                lines[5].split(": ", maxsplit=1)[1],
            ),
        }
        monkeys.append(monkey_dict)

    # start the rounds
    for _ in range(20):
        for i in range(len(monkeys)):
            monkey = monkeys[i]
            for _ in range(len(monkey["items"])):
                counter[monkey["name"]] += 1
                item = monkey["items"].popleft()
                worry = monkey["operation"](item)
                worry = worry // 3
                action = monkey["test"](worry)
                monkeys[parse_action(action)]["items"].append(worry)

    return whats_da_business(counter)


def part_two():
    counter = Counter()
    monkeys = []
    for lines in parse_monkeys():
        monkey_dict = {
            "name": lines[0].strip(":").casefold(),
            "items": deque(
                map(int, lines[1].split(": ", maxsplit=1)[1].split(", "))
            ),
            "operation": parse_ops(lines[2].split(": ", maxsplit=1)[1]),
            "divisor": parse_two_tests(lines[3].split(": ", maxsplit=1)[1]),
            "true": parse_action(lines[4].split(": ", maxsplit=1)[1]),
            "false": parse_action(lines[5].split(": ", maxsplit=1)[1]),
        }
        monkeys.append(monkey_dict)

    _lcm = lcm(*map(lambda m: m["divisor"], monkeys))

    # start the rounds
    for _ in range(1, 10000 + 1):
        for i in range(len(monkeys)):
            monkey = monkeys[i]
            for _ in range(len(monkey["items"])):
                counter[monkey["name"]] += 1
                item = monkey["items"].popleft()
                worry = monkey["operation"](item) % _lcm
                m = (
                    monkey["true"]
                    if worry % monkey["divisor"] == 0
                    else monkey["false"]
                )
                monkeys[m]["items"].append(worry)

    return whats_da_business(counter)


if __name__ == "__main__":
    # print(part_one())
    print(part_two())
