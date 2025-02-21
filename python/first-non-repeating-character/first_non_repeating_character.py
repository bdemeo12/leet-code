from collections import Counter

# Given a string, return the first non-repeating character.


def first_non_repeating_character(someString: str) -> str:

    charCount = Counter(someString)

    for c in charCount:
        if charCount[c] == 1:
            return c

    return ""


if __name__ == "__main__":
    someString = "bababanbabmabak"

    print(first_non_repeating_character(someString))


# def first_non_repeating_character(someString: str) -> str:

#     strings = list(someString)

#     for s in strings:
#         if strings.count(s) == 1:
#             return s

#     return ""
