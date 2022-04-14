# 평균
import sys

count = int(sys.stdin.readline())

valueList = list(map(int, sys.stdin.readline().split()))

maxVal = valueList[0]

for i in valueList:
    if i > maxVal:
        maxVal = i

valueTotal = 0

for i in range(0, count):
    newValue = valueList[i] / maxVal * 100
    valueTotal += newValue

print(valueTotal / count)
