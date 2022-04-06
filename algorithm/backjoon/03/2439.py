# 별찍기 2

import sys

N = int(sys.stdin.readline())
for i in range(1, N + 1):
    empty = " " * (N - i)
    star = "*" * i
    print(f"{empty}{star}")
