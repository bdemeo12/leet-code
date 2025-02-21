# Implement a function to find the length of the longest substring without repeating characters.


def longestSubstring(s: str) -> str:
    left = 0
    char_index = {}
    max_length = 0
    max_substring = ""

    for right in range(len(s)):
        if s[right] in char_index and char_index[s[right]] >= left:
            left = char_index[s[right]] + 1  # Move left pointer to avoid duplicates

        char_index[s[right]] = right  # Store/update index of character
        current_substring = s[left : right + 1]

        if len(current_substring) > max_length:
            max_length = len(current_substring)
            max_substring = current_substring

    return max_substring

    # i, j = 0, len(s) - 1
    # sChars = list(s)
    # uniqueChars = []
    # max = ""
    # while i <= j:
    #     if uniqueChars.count(sChars[i]) == 0:
    #         uniqueChars.append(sChars[i])
    #         i += 1
    #     else:
    #         if len(max) < len(uniqueChars):
    #             max = "".join(uniqueChars)
    #         uniqueChars = []
    # return max


if __name__ == "__main__":
    print(longestSubstring("abcabcbb"))
