from operator import add, mul, sub, truediv

flip = lambda f: lambda *a: f(*reversed(a))

OPS = {
    "+": (add, sub, sub),
    "-": (sub, add, flip(sub)),
    "*": (mul, truediv, truediv),
    "/": (truediv, mul, flip(truediv)),
}


def parse(f):
    vals = {}
    for line in f:
        name, equation = line.split(": ")
        match equation.split():
            case [num]:
                vals[name] = int(num)
            case [a, op, b]:
                vals[name] = a, b, *OPS[op]
    return vals


def compute(vals, i):
    match vals[i]:
        case a, b, f, _, _:
            av, bv = compute(vals, a), compute(vals, b)
            if None in (av, bv):
                return None
            return f(av, bv)
        case _:
            return vals[i]

def part_one(f):
    vals = parse(f)
    return int(compute(vals, "root"))


def part_two(f):
    vals = parse(f)
    vals["humn"] = None
    vals["root"] = *vals["root"][:2], *OPS["-"]

    def solve(i, val):
        match vals[i]:
            case a, b, _, fa, fb:
                match compute(vals, a), compute(vals, b):
                    case av, None:
                        return solve(b, fb(val, av))
                    case None, bv:
                        return solve(a, fa(val, bv))
            case None:
                return val

    return int(solve("root", 0))
