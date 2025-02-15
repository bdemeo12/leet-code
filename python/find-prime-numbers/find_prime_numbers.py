# Write a function that finds all prime numbers up to a given number n.
# 2, 3, 5, 7, 11, 13, 17, 19, 23, 29


def find_primes(n: int) -> list:

    primes = []
    for i in range(n):
        if i == 0 or i == 1 or i == 2:
            primes.append(i)
            continue

        for j in range(1, i + 1):
            if i % j == 0 and j != 1 and j != i:
                # as long as we find one other number that is not i or is not 1, then is it not prime
                break
            else:
                if j == i:
                    primes.append(i)

    return primes


if __name__ == "__main__":
    print(find_primes(30))
