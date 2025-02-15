# Write a function to find the largest number in an array of integers.


def findLargestNumber(l: list) -> int:

    max = 0
    first = True

    for num in l:
        if first or num > max:
            max = num
            first = False

    return max


if __name__ == "__main__":
    l = [1, 22, 400, 2, 0, 343251, 2341, 525, 658, 5453, 153, 67, 33, 5000000]
    print(findLargestNumber(l))
