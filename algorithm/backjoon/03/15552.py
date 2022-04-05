import sys

# 빠른 A+B

T = int(sys.stdin.readline())

lists = []
for i in range(0, T):
    (
        a,
        b,
    ) = map(int, sys.stdin.readline().split())
    lists.append([a, b])

for i in lists:
    print(i[0] + i[1])
