from collections import deque

DIRS = [(1,0,0), (1, -1, 0), (1, 1, 0), (1, 0, -1), (1, 0, 1)]

def add_tuple(a, b):
    return tuple(sum(p) for p in zip(a, b))


def blizzard(board, N, M, t, i, j):
    if (i,j) in ((-1, 0), (N, M -1)):
        return False
    try:
        return (
            board[(i-t) % N, j] == "v"
            or board[(i+t) % N, j] == "^"
            or board[i, (j-t) % M] == ">"
            or board[i, (j+t) % M] == "<"
        )
    except KeyError:
        return True


def shortest_path(board, N, M, begin, end, start_t=0):
    seen = set()
    bfs = deque([(start_t, *begin)])

    while len(bfs) > 0:
        t, i, j = p = bfs.popleft()
        if p in seen:
            continue
        seen.add(p)

        if (i, j) == end:
            return t

        for d in DIRS:
            x = add_tuple(p, d)
            if not blizzard(board, N, M, *x):
                bfs.append(x)


def part_one(f):
    lines = [row[1:-1] for row in f.read().splitlines()[1:-1]]
    board = {(i,j): cell for i, row in enumerate(lines) for j, cell in enumerate(row)}
    N, M = len(lines), len(lines[0])
    board[-1, 0] = board[N, M -1] = "."
    return shortest_path(board, N, M, (-1, 0), (N, M-1))


def part_two(f):
    lines = [row[1:-1] for row in f.read().splitlines()[1:-1]]
    board = {(i,j): cell for i, row in enumerate(lines) for j, cell in enumerate(row)}
    N, M = len(lines), len(lines[0])
    board[-1, 0] = board[N, M -1] = "."

    x = shortest_path(board, N, M, (-1, 0), (N, M-1))
    y = shortest_path(board, N, M, (N, M-1), (-1, 0), x)
    return shortest_path(board, N, M, (-1, 0), (N, M-1), y)
