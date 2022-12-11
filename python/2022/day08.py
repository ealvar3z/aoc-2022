from math import prod

def part_one(f):
    forest = f.read().splitlines()
    total = 0
    for ii, row in enumerate(forest):
        for jj, col in enumerate(row):
            total += (
                all(forest[ii][j] < col for j in range(0, jj))
                or all(forest[ii][j] < col for j in range(jj+1, len(row)))
                or all(forest[i][jj] < col for i in range(0, ii))
                or all(forest[i][jj] < col for i in range(ii+1, len(forest)))
            )
    return total

def part_two(f):
    forest = f.read().splitlines()
    res = []

    for ii, row in enumerate(forest):
        for jj, col in enumerate(row):
            res.append([0, 0, 0, 0])

            for j in range(jj-1, -1, -1):
                res[-1][0] += 1
                if forest[ii][j] >= col:
                    break

            for i in range(ii-1, -1, -1):
                res[-1][1] += 1
                if forest[i][jj] >= col:
                    break

            for j in range(jj+1, len(row)):
                res[-1][2] += 1
                if forest[ii][j] >= col:
                    break

            for i in range(ii+1, len(forest)):
                res[-1][3] += 1
                if forest[i][jj] >= col:
                    break

    return max(prod(x) for x in res)
