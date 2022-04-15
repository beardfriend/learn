import sys

testCasetotal = int(sys.stdin.readline())
lists = []
countList = []
for i in range(0, testCasetotal):
    lists.append(sys.stdin.readline())

for i in lists:
    count = 0
    accumulateNum = 0
    for j in i:
        if j == "O":
            accumulateNum += 1
            count = count + accumulateNum
        if j == "X":
            accumulateNum = 0
    countList.append(count)

for i in countList:
    print(i)
