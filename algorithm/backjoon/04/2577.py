import sys

A = int(sys.stdin.readline())
B = int(sys.stdin.readline())
C = int(sys.stdin.readline())


total = A * B * C


stringTotal = str(total)

totalcheck = {
    "0": 0,
    "1": 0,
    "2": 0,
    "3": 0,
    "4": 0,
    "5": 0,
    "6": 0,
    "7": 0,
    "8": 0,
    "9": 0,
}

for i in range(0, len(stringTotal)):
    totalcheck[stringTotal[i]] += 1


for i in totalcheck:
    print(totalcheck[i])
