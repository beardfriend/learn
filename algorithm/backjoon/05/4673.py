def d(n):
    for i in str(n):
        n += int(i)
    return n


nonSelfNumber = set()

for i in range(1, 10000):
    print(d(i))
    nonSelfNumber.add(d(i))

for j in range(1, 10000):
    if j not in nonSelfNumber:
        print(j)
