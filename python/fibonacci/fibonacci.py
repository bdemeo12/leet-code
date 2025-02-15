# Write a function to generate the Fibonacci sequence up to n numbers.
# the Fibonacci sequence is a sequence in which each element is the sum of the two elements that precede it.
# 0 1 1 2 3 5 8 13 21 34
def fibonacci(n: int) -> list:
    if n == 0:
        return []

    if n == 1:
        return [0]

    if n == 2:
        return [0, 1]

    fib_numbers = [0, 1]

    for i in range(n - 2):
        fib_numbers.append(fib_numbers[-1] + fib_numbers[-2])

    return fib_numbers


if __name__ == "__main__":
    print(fibonacci(100))
