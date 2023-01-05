from collections import deque, namedtuple

N = namedtuple("_", ("id", "val"))

def part_one(f):
    numbers = deque(N(id, int(x)) for id, x in enumerate(f))
    order = list(numbers)

    for x in order:
        idx = next(i for i, y in enumerate(numbers) if x.id == y.id)
        numbers.rotate(-idx)
        numbers.popleft()
        numbers.rotate(-x.val)
        numbers.appendleft(x)

    numbers.rotate(-next(i for i, x in enumerate(numbers) if x.val == 0))
    return numbers[1000 % len(numbers)].val + numbers[2000 % len(numbers)].val + numbers[3000 % len(numbers)].val


def part_two(f):
    nums = deque(N(id, int(x) * 811589153) for id, x in enumerate(f))
    order = list(nums)

    for time in range(10):
        for x in order:
            idx = next(i for i, y in enumerate(nums) if x.id == y.id)
            nums.rotate(-idx)
            nums.popleft()
            nums.rotate(-x.val)
            nums.appendleft(x)

    nums.rotate(-next(i for i, x in enumerate(nums) if x.val == 0))
    return nums[1000 % len(nums)].val + nums[2000 % len(nums)].val + nums[3000 % len(nums)].val
