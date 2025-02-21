# Given a rotated sorted array, implement a function to find a target value.
# Example:
# search_rotated([4,5,6,7,0,1,2], 0)  # Output: 4


def binarySearch(arr: list[int], target: int) -> int:

    left, right = 0, len(arr) - 1

    while left <= right:
        mid = left + right // 2

        if arr[mid] == target:
            return mid

        if arr[left] <= arr[mid]:
            if arr[left] <= target < arr[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if arr[mid] <= target < arr[right]:
                left = mid + 1
            else:
                right = mid - 1

    return -1


if __name__ == "__main__":
    arr = [4, 5, 6, 7, 0, 1, 2]
    print(binarySearch(arr, 7))
