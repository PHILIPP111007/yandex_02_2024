# Description of Yung's diagram:
# http://crei.skoltech.ru/app/data/uploads/sites/42/2020/05/chaban.pdf

n, K = list(map(int, input().split()))

lst = []
for i in range(n):
    lst.append(int(input()))


def func_1(lst) -> int:
    lst = sorted(lst)

    dots = []
    for i in range(n):
        if i == 0:
            for _ in range(lst[i]):
                dots.append(0)
        else:
            raznost = lst[i] - lst[i - 1]
            for _ in range(raznost):
                dots.append(0)
        dots.append(1)

    res = 0
    len_dots = len(dots)
    for i in range(len_dots):
        if dots[i] == 0:
            x = i + K
            if x < len_dots:
                if dots[x] == 1:
                    res += 1

    return res


print(func_1(lst))
