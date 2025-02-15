# Write a function that checks if two strings are anagrams.
# a word, phrase, or name formed by rearranging the letters of another


def isAnagram(s, t) -> bool:
    return sorted(s) == sorted(t)


if __name__ == "__main__":
    print(isAnagram("cinema", "iceman"))
    print(isAnagram("bad", "cat"))
