# A + B -5

plug = True
lists = []
while plug:
    A, B = map(int, input().split())
    if A == 0 and B == 0:
        plug = False
        break
    lists.append(A + B)


for i in lists:
    print(i)
