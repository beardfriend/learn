import sys

X = int(sys.stdin.readline())
count = 0

for i in range(1, X + 1):
    if i < 100:
        count += 1
    else:
        lists = []
        for j in str(i):
            lists.append(int(j))
        if lists[0] - lists[1] == lists[1] - lists[2]:
            count += 1
print(count)
