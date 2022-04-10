# 더하기 사이클
import sys

N = int(sys.stdin.readline())
isChecking = True
count = 0


while isChecking:
    if count != 0:
        first = new // 10
        second = new % 10
    else:
        first = N // 10
        second = N % 10
    new = (second * 10) + ((first + second) % 10)
    if new == N:
        isChecking = False
    count += 1

print(count)
