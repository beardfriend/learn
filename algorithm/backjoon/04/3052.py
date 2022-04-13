# 나머지
import sys

lists = []
for i in range(0, 10):
    num = int(sys.stdin.readline()) % 42
    lists.append(num)

count = 0
for i, num in enumerate(lists):
    cnt = 0
    for j in range(i + 1, len(lists)):

        if num == lists[j]:
            cnt += 1
    if cnt == 0:
        count += 1

print(count)
