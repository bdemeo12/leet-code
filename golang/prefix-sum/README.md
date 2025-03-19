
The prefix sum technique is a way to efficiently calculate the sum of elements within a given range in an array. It involves creating a new array, called the prefix sum array, where each element at index i is the sum of all elements in the original array up to index i. 

## Construction

### To construct the prefix sum array:
1. Initialize a new array prefix_sum with the same size as the input array.
1. Set prefix_sum[0] equal to the original array's first element.
1. Iterate through the original array from index 1 to the end.
1. For each element at index i, calculate prefix_sum[i] as prefix_sum[i-1] + original_array[i].