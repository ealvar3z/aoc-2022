DIGITS = {"2": 2, "1": 1, "0": 0, "-": -1, "=": -2}
ROTATE = {v: k for k, v in DIGITS.items()}

def part_one(f):
    total = 0
    for n in f:
        for i, d in enumerate(n.strip()[::-1]):
            total += DIGITS[d] * 5**i

    res = ""

    while total > 0:
        total, digit = divmod(total, 5)
        if digit > 2:
            digit -= 5
            total += 1
        res += ROTATE[digit]
    
    return res[::-1]

