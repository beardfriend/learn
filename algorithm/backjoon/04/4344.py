# 평균은 넘겠지
import sys

count = int(sys.stdin.readline())

lists = []
for i in range(count):
    lists.append(list(map(int, sys.stdin.readline().split())))


for i in lists:
    total = 0
    percentCount = 0
    for j in range(1, len(i)):
        total += i[j]
    average = total / i[0]
    for t in range(1, len(i)):
        if i[t] > average:
            percentCount += 1
    result = round(percentCount / (len(i) - 1) * 100, 3)
    print(f"{result:.3f}%")
