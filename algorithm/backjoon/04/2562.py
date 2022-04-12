import sys

lists = []
for i in range(0, 9):
    lists.append(int(sys.stdin.readline()))

maxVal = lists[0]
ordering = 1

for i in range(1, 9):
    if lists[i] > maxVal:
        maxVal = lists[i]
        ordering = i + 1

print(maxVal)
print(ordering)
