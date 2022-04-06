# A+B - 8

import sys

T = int(sys.stdin.readline())

lists = []
for i in range(0, T):
    (
        a,
        b,
    ) = map(int, sys.stdin.readline().split())
    lists.append([a, b])

for i in range(0, len(lists)):

    add = lists[i][0] + lists[i][1]
    print("Case #{0}: {1} + {2} = {3}".format(i + 1, lists[i][0], lists[i][1], add))
