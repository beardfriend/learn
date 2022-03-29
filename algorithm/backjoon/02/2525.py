H, M = map(int, input().split())
END = int(input())

H += END // 60
M += END % 60

if M >= 60:
    H += 1
    M -= 60
if H >= 24:
    H -= 24

print(H, M)
