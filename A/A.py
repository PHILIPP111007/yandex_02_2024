from collections import Counter


n = int(input())
a = list(map(int, input().split()))

c = Counter(a)
c = c.items()
m = max(c, key=lambda x: x[1])[1]

sorted_c = sorted(filter(lambda x: x[1] == m, c))
print(sorted_c[-1][0])
