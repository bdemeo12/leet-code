import heapq

# Given an array, find the Kth largest element using a heap.
# Example:
# kth_largest([3,2,1,5,6,4], 2)  # Output: 5


def kth_largest(arr: list[int], k: int) -> int:
    heapq.heapify(arr)

    for _ in range(len(arr) - k):
        heapq.heappop(arr)

    return arr[0]


if __name__ == "__main__":
    arr = [3, 2, 1, 5, 6, 4]

    print(kth_largest(arr, 2))
