n = int(input())
words = [input() for _ in range(n)]


def is_palindrom(s):
    return s == s[::-1]


def func(words) -> set:
    result = set()
    word_to_index = {word: index for index, word in enumerate(words)}
    for i, word in enumerate(words):
        for j in range(len(word) + 1):
            prefix = word[:j]
            suffix = word[j:]
            reversed_suffix = suffix[::-1]
            if (
                is_palindrom(prefix)
                and reversed_suffix in word_to_index
                and i != word_to_index[reversed_suffix]
            ):
                result.add((word_to_index[reversed_suffix] + 1, i + 1))

            if j > 0:
                reversed_prefix = prefix[::-1]
                if (
                    is_palindrom(suffix)
                    and reversed_prefix in word_to_index
                    and i != word_to_index[reversed_prefix]
                ):
                    result.add((i + 1, word_to_index[reversed_prefix] + 1))
    return result


result = func(words)
result = sorted(result, key=lambda x: x[1])
result = sorted(result, key=lambda x: x[0])

for i in result:
    print(*i)
