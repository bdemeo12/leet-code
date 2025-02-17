# Write a function that sums the digits of a number.


def sumAll(n: int) -> int:
    s = str(n)
    l = list(s)

    final = 0
    for tmp in l:
        final += int(tmp)

    return final


print(sumAll(1234))
