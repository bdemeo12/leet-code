# Given a list of words, group anagrams together.
# group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
# # Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]


def findAllAnagrams(anagrams: list[str]) -> list[list[str]]:

    groupedAnagrams = dict()

    for anagram in anagrams:
        a = "".join(sorted(anagram))

        if a in groupedAnagrams:
            groupedAnagrams[a].append(anagram)
        else:
            groupedAnagrams[a] = [anagram]

    return list(groupedAnagrams.values())


if __name__ == "__main__":
    anagrams = ["eat", "tea", "tan", "ate", "nat", "bat"]
    print(findAllAnagrams(anagrams))
