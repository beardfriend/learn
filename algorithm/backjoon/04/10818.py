import sys

N = int(sys.stdin.readline())
lists = list(map(int, sys.stdin.readline().split()))


maxNum = lists[0]
minNum = lists[0]
for j in range(1, len(lists)):
    if lists[j] >= maxNum:
        maxNum = lists[j]
    if lists[j] <= minNum:
        minNum = lists[j]

print(minNum, maxNum)
