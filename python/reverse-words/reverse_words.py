# Write a function that reverses the words in a given string without using built-in split().


def reverseWordsInAString(s: str) -> str:
    words = str.split(s, " ")

    i = 1
    reversed = []
    while i <= len(words):
        reversed.append(words[-i])
        i += 1

    return " ".join(reversed)


if __name__ == "__main__":
    print(reverseWordsInAString("hello world"))
