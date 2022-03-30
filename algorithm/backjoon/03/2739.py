# 구구단
num = input()
for i in range(1, 10):
    print("{0} * {1} = {2}".format(int(num), i, int(num) * i))
