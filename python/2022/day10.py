G = [['?' for _ in range(40)] for _ in range(6)]
p1 = 0

def part_one(f, t=0, x=1):
    data = f.read().strip()
    lines = [x for x in data.split('\n')]
    global p1
    global G
    t1 = t -1
    G[t1//40][t1%40] = ('#' if abs(x-(t1%40))<=1 else ' ')
    if t in [20, 40, 60, 100, 140, 180, 220]:
        p1 += x*t

    for l in lines:
        words = l.split()
        # print(words)
        if words[0] == 'noop':
            t += 1
            part_one(t, x)
        elif words[0] == 'addx':
            t += 1
            part_one(t, x)
            t += 1
            part_one(t, x)
            x += int(words[1])
    print(p1)
    for r in range(6):
        print(''.join(G[r]))
